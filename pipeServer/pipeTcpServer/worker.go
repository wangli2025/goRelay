package pipeserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
)

func worker(conn net.Conn, whiteIpList []string, blackIpList []string) {

	defer conn.Close()

	if pkg.IsBlacklisted(conn.RemoteAddr().String(), blackIpList) {
		goLog.Debug(conn.RemoteAddr().String(), " is black list")
		return
	}

	if !pkg.IsWhitelisted(conn.RemoteAddr().String(), whiteIpList) {
		goLog.Debug(conn.RemoteAddr().String(), " not while list")
		return
	}

	goLog.Info("Client ", conn.RemoteAddr().String(), " is connected.")

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if nil == msg {
			goLog.Error(conn.RemoteAddr(), " RecvMessgae error")
			break
		}

		if string(msg) == "relayConn" {
			relayWorker(conn)
			continue
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

		func() {
			clientConnMapLock.Lock()
			defer clientConnMapLock.Unlock()

			clientConnMap[p.Id] = conn
			goLog.Debug(p.Conn, " has received a client request\nBuf:\n", string(p.Buf), "\nconn:\n", p.Conn, "\nid:", p.Id)
			goLog.Debug("clientConnMaps: ", clientConnMap)
		}()

		func() {
			sendLock.Lock()
			defer sendLock.Unlock()
			pipeprotocol.SendMessage(relayConn, msg)
		}()
	}
}
