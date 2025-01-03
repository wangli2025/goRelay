package main

import (
	"flag"
	"fmt"
	pipeclient_main "goRelay/pipeClient"
	pipeserver_main "goRelay/pipeServer"
	"goRelay/pkg"
	relayclient_main "goRelay/relayClient"
	relayserver_main "goRelay/relayServer"
)

var (
	serverType string
	version    bool
	help       bool
	configFile string
)

func init() {
	flag.StringVar(&serverType, "type", "pipeServer", "server type")
	flag.StringVar(&configFile, "config", "./pipeserver.json", "pipe server conf")
	flag.BoolVar(&version, "version", false, "version")
	flag.BoolVar(&help, "help", false, "help")
}

func main() {
	flag.Parse()

	if version {
		fmt.Println("version:", pkg.Version, "buildAt:", pkg.BuildAt, "gitCommit:", pkg.GitCommit)
		return
	}
	if help {
		flag.Usage()
		return
	}

	fmt.Println("server Type:", serverType)
	fmt.Println("server config path:", configFile)
	fmt.Println("version:", pkg.Version)
	fmt.Println("buildAt:", pkg.BuildAt)
	fmt.Println("gitCommit:", pkg.GitCommit)

	switch serverType {
	case "relayServer":
		relayserver_main.RunRelayServer(configFile)

	case "pipeServer":
		pipeserver_main.RunPipeServer(configFile)

	case "relayClient":
		relayclient_main.RunRelayClientServer(configFile)

	case "pipeClient":
		pipeclient_main.RunPipeClientServer(configFile)

	default:
		fmt.Println("server type error,not fount ", serverType)
	}
}
