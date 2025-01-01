package main

import "flag"

var (
	configFile      string
	pipeServerAddr  string
	relayClientAddr string
	debugLog        bool
	version         bool
	help            bool
)

type Config struct {
	PipeServerAddr  string `json:"pipe_server_addr"`
	RelayClientAddr string `json:"relay_client_addr"`
	DebugLog        bool   `json:"debug_log"`
}

func init() {
	flag.StringVar(&configFile, "config", "./pipeserver.json", "pipe server conf")
	flag.StringVar(&pipeServerAddr, "pipeServerAddr", "127.0.0.1:8888", "pipe server addr")
	flag.StringVar(&relayClientAddr, "relayClientAddr", "127.0.0.1:10011", "relay client addr")
	flag.BoolVar(&debugLog, "debug", false, "debug log")
}
