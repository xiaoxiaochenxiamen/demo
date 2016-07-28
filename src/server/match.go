package server

import (
	"proto_go"
	"proto_pb"
	"sort"
	"sync"
	"time"
	"utils"
)

const (
	MatchScanf    = 2000 * time.Millisecond
	TeamPlayerMax = 5
	MatchMin      = TeamPlayerMax * 2
	MatchScoreMax = 50
)

var TeamOne = []int{8, 7, 4, 3, 1}
var TeamTwo = []int{9, 6, 5, 2, 0}

func matchTimeScanf() {
	tick := time.NewTicker(MatchScanf)
	for {
		<-tick.C
		MatchServer.rw.RLock()
		playerNum := len(MatchServer.player)
		if playerNum >= MatchMin {
			byScore := make([]*PlayerMatch, playerNum)
			pos := 0
			for id, score := range MatchServer.player {
				byScore[pos] = newPlayerMatch(id, score)
				pos++
			}
			MatchServer.rw.RUnlock()
			sort.Sort(ByScore(byScore))
			matchTeam(byScore)
		} else {
			MatchServer.rw.RUnlock()
		}
	}
}

type MatchPool struct {
	rw     *sync.RWMutex
	player map[int32]int32
}

func newMatchPool() *MatchPool {
	return &MatchPool{
		rw:     &sync.RWMutex{},
		player: make(map[int32]int32),
	}
}

func (m *MatchPool) insert(id int32, score int32) {
	m.rw.Lock()
	m.player[id] = score
	m.rw.Unlock()
}

func (m *MatchPool) len() int {
	m.rw.Lock()
	l := len(m.player)
	m.rw.Unlock()
	return l
}

var MatchServer = newMatchPool()

func handleRequsetMatch(packet *ClientPacket) {
	requst, err := proto_pb.ReadTestPBRequstMatch(packet.buff)
	if nil == err {
		requsetMatch(requst.GetUId(), requst.GetScore())
	}
}

func requsetMatch(userId int32, score int32) {
	MatchServer.insert(userId, score)
}

type PlayerMatch struct {
	UserId int32
	Score  int32
}

func newPlayerMatch(id int32, score int32) *PlayerMatch {
	return &PlayerMatch{
		UserId: id,
		Score:  score,
	}
}

type ByScore []*PlayerMatch

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func matchTeam(byScore []*PlayerMatch) {
	l := len(byScore)
	for i := 0; i < l; {
		if byScore[i+MatchMin-1].Score-byScore[i].Score > MatchScoreMax {
			i++
		} else {
			packetTeam(byScore[i : i+MatchMin])
			i = i + MatchMin
		}
	}
}

func packetTeam(player []*PlayerMatch) {
	team1 := make([]*testPB.Player, TeamPlayerMax)
	team2 := make([]*testPB.Player, TeamPlayerMax)
	for i := 0; i < TeamPlayerMax; i++ {
		team1[i] = &testPB.Player{
			UId:   &(player[TeamOne[i]].UserId),
			Score: &(player[TeamOne[i]].Score),
		}
		team2[i] = &testPB.Player{
			UId:   &(player[TeamTwo[i]].UserId),
			Score: &(player[TeamTwo[i]].Score),
		}
	}
	team1Win := (0 == utils.Rand()%2)
	result := &testPB.MatchResult{
		IsTeam1Win: &team1Win,
		Team1:      team1,
		Team2:      team2,
	}
	data, err := proto_pb.WriteTestPBMatchResult(result)
	if nil == err {
		sendMessageToClient(data)
	}
}
