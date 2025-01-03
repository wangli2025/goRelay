package relayclient_main

import (
	"fmt"
	"goRelay/pkg"
	relaytcpclient "goRelay/relayClient/relayTcpClient"
)

func RunRelayClientServer(configFile string) {

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

	var realServerInfoMap map[string]string
	realServerInfoMap = make(map[string]string, 0)
	for _, v := range config.RealServerInfo {
		realServerInfoMap[v.ID] = v.RealServerAddr
	}
	relaytcpclient.RunTcpServer(config.ListenRelayClientAddr, config.WhiteIpList, realServerInfoMap)

}
