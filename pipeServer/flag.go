package main

import "flag"

var (
	configFile string
	version    bool
	help       bool
)

type Config struct {
	ListenPipeServerAddr string   `json:"listen_pipe_server_addr"`
	WhiteIpList          []string `json:"white_ip_list"`
	BlackIpList          []string `json:"black_ip_list"`
	DebugLog             bool     `json:"debug_log"`
}

func init() {
	flag.StringVar(&configFile, "config", "./pipeserver.json", "pipe server conf")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
}
