package main

import (
	"flag"
	"fmt"
	pipetcpclient "goRelay/pipeClient/pipeTcpClient"
	"goRelay/pkg"
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

	go pipetcpclient.ConnectToRelay(relayClientAddr)
	pipetcpclient.ConnectToPipeServer(pipeServerAddr)
}
