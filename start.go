package main

import (
	"github.com/jessevdk/go-flags"
	"go-proxy/parser"
	"go-proxy/proxy"
	"os"
	"os/signal"
	"syscall"
)

var (
	opts   Option
	isHold = false
)

func init() {
	flags.Parse(&opts)
}

type Option struct {
	Proxy   bool   `short:"P" long:"proxy" description:"开启代理"`
	Port    int    `short:"p" long:"port" description:"代理连接端口" default:"8989"`
	Check   bool   `long:"check" description:"查看网卡名字和IP"`
	Open    bool   `short:"o" long:"open" description:"打开网卡监听"`
	IfcName string `short:"i" long:"interface" description:"监听网卡名字" default:"eth0"`
	TCPPort int    `short:"t" long:"tcpPort" description:"监听tcp端口" default:"80"`
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	/**
	go func() {
		<-c
		os.Exit(0)
	}()
	**/

	// 打印网卡信息
	if opts.Check {
		parser.ParseAndPrint()
	}

	// 启动程序
	if opts.Proxy {
		go proxy.HttpProxy(opts.Port)
		isHold = true
	}

	// 打开网卡监听
	if opts.Open {
		go parser.ParsePacket(opts.IfcName, opts.TCPPort)
		isHold = true
	}

	if !isHold {
		c <- syscall.SIGQUIT
	}

	<-c
	os.Exit(0)
}
