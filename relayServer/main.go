package main

import (
	"flag"
	"fmt"
	"goRelay/pkg"
	relaytcpserver "goRelay/relayServer/relayTcpServer"
)

func main() {

	flag.Parse()

	if version {
		fmt.Println("version:", pkg.Version, "gitCommit:", pkg.GitCommit)
		return
	}
	if help {
		flag.Usage()
		return
	}

	var config Config
	if pkg.LoadConfig(configFile, &config) != nil {
		fmt.Println("read config file error")
		return
	}
	fmt.Println("config:", config)

	goLog := pkg.NewLogger()
	if config.DebugLog {
		goLog.SetLogger(pkg.DebugLevel)
	} else {
		goLog.SetLogger(pkg.LogLevel)
	}
	go relaytcpserver.ConnectPipeServer(config.PipeServerAddr)

	relaytcpserver.RunTcpServer(config.ListenRelayServerAddr, config.WhiteIpList, config.Id)
}
