package main

import "flag"

var (
	realServer string
	listenAddr string
	version    bool
	help       bool
	debugLog   bool
)

func init() {
	flag.StringVar(&realServer, "realServer", "127.0.0.1:80", "real server")
	flag.StringVar(&listenAddr, "listenAddr", ":10011", "listen addr")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&debugLog, "debug", false, "debug log")
}
