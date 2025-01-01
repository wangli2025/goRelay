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

	go pipetcpclient.ConnectToRelay(config.RelayClientAddr)
	pipetcpclient.ConnectToPipeServer(config.PipeServerAddr)
}
