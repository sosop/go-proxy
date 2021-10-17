package proxy

import (
	"fmt"
	"log"
	"net"
)

func HttpProxy(port int) {
	fmt.Println("代理端口：", port)
	ln, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		log.Fatalf("监听错误：%s\n", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("链接错误：%s\n", err.Error())
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// TODO
}
