package main

import (
	"fmt"
	"net"
	_ "sync"
)

type ThreadPool struct {
	workers int
	msgList chan net.Conn
	term    bool
}

func createThreadPool(w int) ThreadPool {
	tp := ThreadPool{
		workers: w,
		msgList: make(chan net.Conn, w),
		term:    false}

	for i := 0; i < tp.workers; i++ {
		go tp.handleEcho()
	}
	return tp
}

func (tp *ThreadPool) processConnection(c net.Conn) {
	tp.msgList <- c
}

func (tp *ThreadPool) handleEcho() {
	for !tp.term {
		conn := <-tp.msgList
		var inmsg = make([]byte, 255)

		_, e := conn.Read(inmsg)
		if e != nil {
			fmt.Println(e.Error())
		} else {
			conn.Write([]byte(inmsg))
		}

		conn.Close()
	}
}
