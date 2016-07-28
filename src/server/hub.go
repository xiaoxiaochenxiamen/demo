package server

import (
	"proto_go"
)

var HubChan = NewHubChan()

func NewHubChan() chan *ClientPacket {
	return make(chan *ClientPacket, 1024)
}

func InitHub() {
	go clientHub()
}

func clientHub() {
	for {
		packet := <-HubChan
		switch packet.Id {
		case testPB.MsgId_REQUEST_MATCH:
			handleRequsetMatch(packet)
		default:
		}
	}
}

type ClientPacket struct {
	Id   testPB.MsgId
	Buff []byte
}

func NewClientPacket(id int32, data []byte) *ClientPacket {
	return &ClientPacket{
		Id:   testPB.MsgId(id),
		Buff: data,
	}
}
