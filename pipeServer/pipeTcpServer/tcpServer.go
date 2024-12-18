package pipeserver

import (
	"net"
)

func ListenTcpServer(pipeServerListen string) {
	listen, err := net.Listen("tcp", pipeServerListen)
	if err != nil {
		goLog.Error("listen error", err)
		return
	}
	goLog.Info("listen", pipeServerListen, " successful")

	for {
		conn, err := listen.Accept()
		if err != nil {
			goLog.Error("accept error", err)
			continue
		}
		go worker(conn)
	}
}
