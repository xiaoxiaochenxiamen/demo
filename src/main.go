package main

import (
	"client"
	"server"
	"time"
)

func main() {
	go server.StartServer()
	time.Sleep(1 * time.Second)
	client.StarClient()
}
