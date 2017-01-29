package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", ":9001")

	tp := createThreadPool(3)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	for true {
		incomming, _ := conn.Accept()
		tp.processConnection(incomming)
	}

}
