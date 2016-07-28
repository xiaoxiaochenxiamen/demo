package server

import (
	"bytes"
	"fmt"
	"net"
	"proto_pb"
	"utils"
)

const (
	ConReadBuffSize = 256
	ConBuffSize     = 1024
	Net             = "tcp"
)

var ClientCon *net.TCPConn

func StartServer() {
	InitHub()
	InitMatch()
	listen, ok := listenTcp()
	if ok {
		accept(listen)
	}
}

func listenTcp() (*net.TCPListener, bool) {
	netAddr := Config.GetServerAddress()
	tcpAddr, err := net.ResolveTCPAddr(Net, netAddr)
	if nil != err {
		fmt.Println("TcpAddr err : ", err.Error())
		return nil, false
	} else {
		listen, err := net.ListenTCP(Net, tcpAddr)
		if nil != err {
			fmt.Println("ListenTCP err : ", err.Error())
			return nil, false
		} else {
			fmt.Println("listen ", Net, netAddr)
			return listen, true
		}
	}
}

func accept(listen *net.TCPListener) {
	for {
		tcpCon, err := listen.AcceptTCP()
		if nil == err {
			tcpCon.SetKeepAlive(true)
			ClientCon = tcpCon
			go HandleCon(tcpCon, HubChan)
		}
	}
}

type ConBuff struct {
	expect     int
	packeId    int32
	isReadNext bool
	con        *net.TCPConn
	*bytes.Buffer
}

func (c *ConBuff) hasPacket() bool {
	if c.isReadNext {
		return c.expect <= c.Len()
	} else {
		return proto_pb.PacketHeadLen <= c.Len()
	}
}

func NewConBuff() *ConBuff {
	conBuff := &ConBuff{
		expect:     8,
		isReadNext: false,
		Buffer:     &bytes.Buffer{},
	}
	return conBuff
}

func HandleCon(con *net.TCPConn, hub chan *ClientPacket) {
	buff := NewConBuff()
	for {
		temp := make([]byte, ConReadBuffSize)
		n, err := con.Read(temp)
		if nil == err && n > 0 {
			if ConReadBuffSize == n {
				buff.Write(temp)
			} else {
				buff.Write(temp[:n])
			}
			if buff.hasPacket() {
				buff.readPacket(hub)
			}
		}
	}
}

func (c *ConBuff) readPacket(hub chan *ClientPacket) {
	for c.hasPacket() {
		if c.isReadNext {
			c.readPacketData(hub)
		} else {
			c.readPacketHead()
		}
	}
}

func (c *ConBuff) readPacketHead() {
	expectCode := make([]byte, 4)
	c.Read(expectCode)
	c.expect = int(utils.DecodeInt32(expectCode)) - proto_pb.PacketHeadLen
	packetIdCode := make([]byte, 4)
	c.Read(packetIdCode)
	c.packeId = utils.DecodeInt32(packetIdCode)
	c.isReadNext = true
}

func (c *ConBuff) readPacketData(hub chan *ClientPacket) {
	data := make([]byte, c.expect)
	c.Read(data)
	c.isReadNext = false
	packet := NewClientPacket(c.packeId, data)
	hub <- packet
}

func sendMessageToClient(msg []byte) {
	ClientCon.Write(msg)
}
