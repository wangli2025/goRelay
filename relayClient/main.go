package main

import (
	"flag"
	"fmt"
	"goRelay/pkg"
	relaytcpclient "goRelay/relayClient/relayTcpClient"
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

	go relaytcpclient.DeleteHistoryClientMessage()

	relaytcpclient.RunTcpServer(listenAddr, realServer)

}
