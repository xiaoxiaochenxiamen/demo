package server

import (
	"proto_go"
)

var HubChan = newHubChan()

func newHubChan() chan *ClientPacket {
	return make(chan *ClientPacket, 1024)
}

func InitHub() {
	go clientHub()
}

func clientHub() {
	for {
		packet := <-HubChan
		switch packet.id {
		case testPB.MsgId_REQUEST_MATCH:
			handleRequsetMatch(packet)
		default:
		}
	}
}

type ClientPacket struct {
	id   testPB.MsgId
	buff []byte
}

func newClientPacket(id int32, data []byte) *ClientPacket {
	return &ClientPacket{
		id:   testPB.MsgId(id),
		buff: data,
	}
}
