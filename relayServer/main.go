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

	goLog := pkg.NewLogger()
	if debugLog {
		goLog.SetLogger(pkg.DebugLevel)
	} else {
		goLog.SetLogger(pkg.LogLevel)
	}

	go relaytcpserver.DeleteHistoryClientMessage()

	go relaytcpserver.ConnectPipeServer(pipeServerAddr)

	relaytcpserver.RunTcpServer(listenAddr)

}
