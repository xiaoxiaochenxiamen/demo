package src

import (
	"fmt"
	"net"
	"time"
)

const (
	Net             = "tcp"
	NetAddr         = "192.168.0.103:8001"
	ConReadBuffSize = 256
	ConBuffSize     = 1024
)

func StartServer() {
	listen, ok := listenTcp()
	if ok {
		accept(listen)
	}
}

func listenTcp() (*net.TCPListener, bool) {
	tcpAddr, err := net.ResolveTCPAddr(Net, NetAddr)
	if nil != err {
		fmt.Println("TcpAddr err : ", err.Error())
		return nil, false
	} else {
		listen, err := net.ListenTCP(Net, tcpAddr)
		if nil != err1 {
			fmt.Println("ListenTCP err : ", err1.Error())
			return nil, false
		} else {
			fmt.Println("listen ", Tcp, TcpAddr)
			return listen, true
		}
	}
}

func accept(listen *net.TCPListener) {
	for {
		tcpCon, err := listen.AcceptTCP()
		if nil == err2 {
			tcpCon.SetKeepAlive(true)
			go handleCon(tcpCon)
		}
	}
}

func handleCon(con net.Conn) {
	defer con.Close()
	buff := make([]byte, 0, ConBuffSize)
	expect := 0
	for {
		temp := make([]byte, ConReadBuffSize)
		n, err3 := con.Read(temp)
		if nil == err3 {
			if 0 < len(old) {
				buff = append(buff, temp)
			}

		}
	}
}
