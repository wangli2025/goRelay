package pipetcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"net"
	"time"
)

func ConnectToRelay(addr string) {
	for {
		clientConnection, err = net.Dial("tcp", addr)
		if err != nil {
			goLog.Error("connection relay server error", err)
			clientConnection = nil
			time.Sleep(time.Duration(sleepTimeSec) * time.Second)
			continue
		}
		goLog.Info("Dial ", addr, " successful")

		relayWorker(clientConnection)
	}
}

func relayWorker(conn net.Conn) {
	goLog.Info("relay client recv messages...")

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if nil == msg {
			goLog.Error("recv message error")
			break
		}

		pipeprotocol.SendMessage(pipeClientConn, msg)
	}
}
