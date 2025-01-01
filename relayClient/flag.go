package main

import "flag"

var (
	configFile string
	version    bool
	help       bool
)

//	type Config struct {
//		Id                    string   `json:"id"`
//		RealServerAddr        string   `json:"real_Server_Addr"`
//		ListenRelayClientAddr string   `json:"listen_relay_client_addr"`
//		WhiteIpList           []string `json:"white_ip_list"`
//		DebugLog              bool     `json:"debug_log"`
//	}

type Config struct {
	ListenRelayClientAddr string       `json:"listen_relay_client_addr"`
	WhiteIpList           []string     `json:"white_ip_list"`
	DebugLog              bool         `json:"debug_log"`
	RealServerInfo        []RealServer `json:"realServerInfo"`
}

type RealServer struct {
	ID             string `json:"id"`
	RealServerAddr string `json:"real_Server_Addr"`
}

func init() {
	flag.StringVar(&configFile, "config", "./pipeserver.json", "pipe server conf")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
}
