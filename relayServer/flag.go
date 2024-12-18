package main

import "flag"

var (
	pipeServerAddr string
	listenAddr     string
	version        bool
	help           bool
	debugLog       bool
)

func init() {
	flag.StringVar(&pipeServerAddr, "pipeAddr", "127.0.0.1:8888", "pipe addr")
	flag.StringVar(&listenAddr, "listenAddr", ":10010", "listen addr")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&debugLog, "debug", false, "debug log")
}
