package main

import (
	"fmt"
	"net"

	"github.com/mbidewell/go_demos/threadpool"
)

type EchoConn struct {
	conn net.Conn
}

func main() {
	conn, err := net.Listen("tcp", ":9001")

	tp := threadpool.CreateThreadPool(3)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()

	for true {
		incomming, _ := conn.Accept()
		cw := EchoConn{conn: incomming}
		tp.SubmitWork(cw)
	}

}

func (ec EchoConn) Execute() {
	var inmsg = make([]byte, 255)

	_, e := ec.conn.Read(inmsg)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		ec.conn.Write([]byte(inmsg))
	}

	ec.conn.Close()
}
