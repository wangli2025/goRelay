package relayserver_main

import (
	"fmt"
	"goRelay/pkg"
	relaytcpserver "goRelay/relayServer/relayTcpServer"
)

func RunRelayServer(configFile string) {

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
