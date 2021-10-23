package proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func HttpProxy(port int) {
	log.Println("代理端口：", port)
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
	if conn == nil {
		return
	}
	defer conn.Close()
	log.Printf("remote address: %v\n", conn.RemoteAddr())

	var buf [1024*9]byte

	n, err := conn.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, URL, target string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &URL)
	hostPortURL, err := url.Parse(URL)
	if err != nil {
		log.Println(err)
		return
	}

	if method == "CONNECT" {
		target = hostPortURL.Scheme + ":" + hostPortURL.Opaque
	} else {
		target = hostPortURL.Host
		if strings.Index(hostPortURL.Host, ":") == -1 {
			target = hostPortURL.Host + ":80"
		}
	}

	server, err := net.Dial("tcp", target)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}

	go io.Copy(server, conn)
	io.Copy(conn, server)
}
