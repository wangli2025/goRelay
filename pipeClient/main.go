package pipeclient_main

import (
	"fmt"
	pipetcpclient "goRelay/pipeClient/pipeTcpClient"
	"goRelay/pkg"
)

func RunPipeClientServer(configFile string) {

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
