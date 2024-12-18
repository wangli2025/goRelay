package main

import "flag"

var (
	pipeServerAddress string
	version           bool
	help              bool
	debugLog          bool
)

func init() {
	flag.StringVar(&pipeServerAddress, "listen", ":8888", "pipe server listen")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&debugLog, "debug", false, "debug log")
}
