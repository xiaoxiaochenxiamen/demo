package client

import (
	"fmt"
	"net"
	"time"
)

func client(n int) {
	// tcpAddr := &net.TCPAddr{
	//  IP:   "127.0.0.1",
	//  Port: 8001,
	// }
	con, err1 := net.Dial(Tcp, TcpAddr)
	defer con.Close()
	if nil != err1 {
		fmt.Println("Dial err : ", err1.Error())
	}
	fmt.Println("net dail ok")
	for i := 0; i < 10; i++ {
		_, err := con.Write([]byte{byte(n + i)})
		if nil != err {
			fmt.Println(err)
		}
	}
}
