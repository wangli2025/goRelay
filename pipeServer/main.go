package pipeserver_main

import (
	"fmt"
	pipeserver "goRelay/pipeServer/pipeTcpServer"
	"goRelay/pkg"
)

func RunPipeServer(configFile string) {

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
