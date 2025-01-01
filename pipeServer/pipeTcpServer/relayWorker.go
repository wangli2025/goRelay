package pipeserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
)

func relayWorker(conn net.Conn) {
	relayConn = conn
	pipeprotocol.SendMessage(conn, []byte("isok"))
	defer func() {
		relayConn = nil
	}()

	goLog.Info(conn.RemoteAddr().String(), " is relay worker...")

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if msg == nil {
			goLog.Error(conn.RemoteAddr(), " RecvMessgae error")
			break
		}

		var p pipeprotocol.ClientProtocolInfo

		func() {
			clientConnMapLock.Lock()
			defer clientConnMapLock.Unlock()
			clientConnMap[p.Id] = conn
		}()

		err := pkg.JsonUnmarshal(msg, &p)
		if err != nil {
			goLog.Error("Error: JSON unmarshal failed:", err, " jsonData:", msg)
			break
		}

		goLog.Debug("p.id: ", p.Id)

		relayServerConn, isok := func() (net.Conn, bool) {
			clientConnMapLock.Lock()
			defer clientConnMapLock.Unlock()
			r, e := clientConnMap[p.Id]
			return r, e
		}()

		if !isok {
			goLog.Error("Error: not fount relayServer Conn")
			continue
		}
		goLog.Debug("relayServerConn: ", relayServerConn)
		goLog.Debug(conn.RemoteAddr().String(), " RecvMessgaes...")

		pipeprotocol.SendMessage(relayServerConn, msg)
	}
}
