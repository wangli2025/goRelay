package main

import "flag"

var (
	version    bool
	help       bool
	configFile string
)

type Config struct {
	Id                    string   `json:"id"`
	PipeServerAddr        string   `json:"pipe_server_addr"`
	ListenRelayServerAddr string   `json:"listen_relay_server_addr"`
	WhiteIpList           []string `json:"white_ip_list"`
	DebugLog              bool     `json:"debug_log"`
}

func init() {
	flag.StringVar(&configFile, "config", "./pipeserver.json", "pipe server conf")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
}
