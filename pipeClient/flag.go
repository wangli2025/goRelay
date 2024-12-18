package main

import "flag"

var (
	pipeServerAddr  string
	relayClientAddr string
	debugLog        bool
	version         bool
	help            bool
)

func init() {
	flag.StringVar(&pipeServerAddr, "pipeServerAddr", "127.0.0.1:8888", "pipe server addr")
	flag.StringVar(&relayClientAddr, "relayClientAddr", "127.0.0.1:10011", "relay client addr")
	flag.BoolVar(&debugLog, "debug", false, "debug log")
}
