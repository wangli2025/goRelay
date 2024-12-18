package pipeserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"net"
)

func relayWorker(conn net.Conn) {
	relayConn = conn
	pipeprotocol.SendMessage(conn, []byte("isok"))
	defer func() {
		relayConn = nil
	}()

	goLog.Info(conn.RemoteAddr().String(), "is relay worker...")

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if msg == nil {
			goLog.Error("RecvMessgae error")
			break
		}

		goLog.Debug(conn.RemoteAddr().String(), " RecvMessgaes...")

		pipeprotocol.SendMessage(clientConn, msg)
	}
}
