package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"proto_go"
	"proto_pb"
	"server"
	"strconv"
	"strings"
	"utils"
)

const (
	MaxClient = 2000
)

var ClinetTcp *net.TCPConn
var ClientHubChan = server.NewHubChan()

type PlayerManager struct {
	player map[int32]int32
}

var AllPlayer = newAllPlayer()

func newAllPlayer() *PlayerManager {
	return &PlayerManager{
		player: make(map[int32]int32),
	}
}

func (p *PlayerManager) insert(id int32, score int32) {
	p.player[id] = score
}

func tcpConnet() {
	addr := server.Config.GetServerAddress()
	tcpAddr, err := net.ResolveTCPAddr(server.Net, addr)
	if nil != err {
		panic(err)
	}
	tcpCon, err := net.DialTCP(server.Net, nil, tcpAddr)
	if nil != err {
		panic(err)
	}
	ClinetTcp = tcpCon
	go server.HandleCon(tcpCon, ClientHubChan)
	go clientHub()

}

func clientHub() {
	for {
		packet := <-ClientHubChan
		switch packet.Id {
		case testPB.MsgId_MATCH_RESULT:
			handleMatchResult(packet)
		default:
		}
	}
}

func handleMatchResult(packet *server.ClientPacket) {
	result, err := proto_pb.ReadTestPBMatchResult(packet.Buff)
	if nil == err {
		fmt.Println(result.GetIsTeam1Win(), result.GetTeam1(), result.GetTeam2())
	}
}

func generateParams(stdParams string) []interface{} {
	newParams := make([]interface{}, 0)
	for _, param := range strings.Split(stdParams, ",") {
		if strings.Contains(param, `"`) {
			newParams = append(newParams, param)
		} else {
			newParamsInt, err := strconv.Atoi(param)
			if err == nil {
				newParams = append(newParams, newParamsInt)
			}

		}
	}
	return newParams
}

func StdLines() <-chan string {
	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	go func() {
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading standard input: ", err)
		}
	}()

	return ch
}

func StarClient() {
	tcpConnet()
	fmt.Println("开始启动测试")
	printHelp()
	stdIn := StdLines()
	pos := 0
	for {
		line := <-stdIn
		fmt.Printf("[In] %v\n", line)
		info := strings.Split(line, " ")

		switch info[0] {

		case "client":
			pos = cmdPase(pos, info)
		case "push":
			pushAllUser()
		case "exit":
			fmt.Println("exit ok")
			return
		case "help":
			printHelp()
		default:
			fmt.Println("无效参数")
			printHelp()
		}
	}
}

func pushAllUser() {
	for userId, score := range AllPlayer.player {
		requstMatch(userId, score)
	}
}

func cmdPase(pos int, info []string) int {
	switch len(info) {
	case 2:
		pos = cmdPaseClient2(pos, info)
	case 3:
		pos = cmdPaseClient3(pos, info)
	default:
		printHelp()
	}
	return pos
}

func cmdPaseClient2(pos int, info []string) int {
	num, err := strconv.Atoi(info[1])
	if nil == err {
		if num+pos > MaxClient {
			fmt.Println("超过最大测试数量: ", MaxClient)
		} else {
			addClient(int32(pos), int32(pos+num))
			pos = pos + num
		}
	} else {
		printHelp()
	}
	return pos
}
func cmdPaseClient3(pos int, info []string) int {
	uid, err := strconv.Atoi(info[1])
	if nil == err {
		score, err := strconv.Atoi(info[2])
		if nil == err {
			requstMatch(int32(uid), int32(score))
			pos = pos + 1
		} else {
			printHelp()
		}
	} else {
		printHelp()
	}
	return pos
}

func printHelp() {
	fmt.Println("键入 client [number](空格隔开)            例: client 100   ------ 新增100个测试用户")
	fmt.Println("键入 client [UserId] [Score](空格隔开)    例: client 1 20  ------ 新增/设置 单个用户积分")
	fmt.Println("键入 push            ------------------------------------------ 全部测试用户匹配")
	fmt.Println("键入 exit            ------------------------------------------ 退出测试")
}

func addClient(start int32, end int32) {
	for i := start; i < end; i++ {
		score := int32(utils.Rand() % 200)
		AllPlayer.insert(i, score)
		requstMatch(i, score)
	}
}

func sendMsgToServer(buff []byte) {
	ClinetTcp.Write(buff)
}

func requstMatch(id int32, score int32) {
	buff, err := proto_pb.WriteTestPBRequstMatch(id, score)
	if nil == err {
		sendMsgToServer(buff)
	}
}
