package main

import (
	"flag"
	"fmt"
	pipeserver "goRelay/pipeServer/pipeTcpServer"
	"goRelay/pkg"
)

func main() {
	flag.Parse()

	if version {
		fmt.Println("version:", pkg.Version, "buildAt:", pkg.BuildAt, "gitCommit:", pkg.GitCommit)
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

	pipeserver.ListenTcpServer(config.ListenPipeServerAddr, config.WhiteIpList, config.BlackIpList)

}
