package pipeserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
)

func worker(conn net.Conn) {
	goLog.Info("Client", conn.RemoteAddr().String(), "is connected.")

	defer conn.Close()

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if nil == msg {
			goLog.Error("RecvMessgae error")
			break
		}

		if string(msg) == "relayConn" {
			relayWorker(conn)
		} else {
			clientConn = conn
		}

		if nil == relayConn {
			goLog.Error("relayConn is nil...")
			break
		}
		goLog.Debug(conn.RemoteAddr().String(), "is client worker...")

		var p pipeprotocol.ClientProtocolInfo
		err := pkg.JsonUnmarshal(msg, &p)
		if err != nil {
			goLog.Error("Error: JSON unmarshal failed:", err, " jsonData:", msg)
			break
		}
		goLog.Debug(p.Conn, " has received a client request\nBuf:\n", string(p.Buf), "\nconn:\n", p.Conn)

		func() {
			sendLock.Lock()
			defer sendLock.Unlock()
			pipeprotocol.SendMessage(relayConn, msg)
		}()
	}
}
