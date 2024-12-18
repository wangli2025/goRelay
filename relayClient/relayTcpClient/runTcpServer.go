package relaytcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"log"
	"net"
	"time"
)

func RunTcpServer(addr string, realServerAddr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("listen error", err)
		return
	}
	goLog.Info("listen ", addr, " successful")

	for {
		relayConn, err = listen.Accept()
		if err != nil {
			goLog.Error("accept error", err)
		}

		worker(relayConn, realServerAddr)
	}
}

func worker(conn net.Conn, realServerAddr string) {
	goLog.Info(conn.RemoteAddr().String(), " is connecttion")

	for {
		msg := pipeprotocol.RecvMessgae(conn)
		if msg == nil {
			goLog.Error("recv relay server error")
			break
		}

		var p pipeprotocol.ClientProtocolInfo
		err := pkg.JsonUnmarshal(msg, &p)
		if err != nil {
			goLog.Error("json unmarshal error ", err)
			continue
		}
		goLog.Debug(p.Conn, " The message has been received.")

		_, isok := clientConnections[p.Conn]
		if !isok {
			realServer, err := net.Dial("tcp", realServerAddr)
			if err != nil {
				goLog.Error("dial real server error", err)
				break
			}

			go connectionRealServer(realServer, p.Conn)
			var c pipeprotocol.ClientInfo
			c.Conn = realServer
			c.Time = time.Now().Unix()

			func() {
				clientMapMutex.Lock()
				defer clientMapMutex.Unlock()
				clientConnections[p.Conn] = c
			}()

		}
		cInfo, _ := clientConnections[p.Conn]

		sendLen := 0
		for sendLen < len(p.Buf) {
			n, err := cInfo.Conn.Write(p.Buf[sendLen:])
			if err != nil {
				log.Println("to real server write error", err)
				break
			}
			sendLen += n
		}
		cInfo.Time = time.Now().Unix()

		func() {
			clientMapMutex.Lock()
			defer clientMapMutex.Unlock()
			clientConnections[p.Conn] = cInfo
		}()
	}
}
