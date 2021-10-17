package main

import (
	"github.com/jessevdk/go-flags"
	"go-proxy/proxy"
	"os"
	"os/signal"
	"syscall"
)

var (
	opts Option
)

func init() {
	flags.Parse(&opts)
}

type Option struct {
	Proxy bool `short:"P" long:"proxy" description:"开启代理"`
	Port  int  `short:"p" long:"port" description:"代理连接端口" default:"8989"`
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c
		os.Exit(0)
	}()
	// 启动程序
	if opts.Proxy {
		proxy.HttpProxy(opts.Port)
	}
}
