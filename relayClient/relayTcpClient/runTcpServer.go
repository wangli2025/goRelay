package relaytcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"log"
	"net"
	"time"
)

func RunTcpServer(addr string, whitelist []string, realServerInfoMap map[string]string) {
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

		worker(relayConn, whitelist, realServerInfoMap)
	}
}

func worker(conn net.Conn, whitelist []string, realServerInfoMap map[string]string) {
	goLog.Info(conn.RemoteAddr().String(), " is connecttion")

	defer conn.Close()

	if !pkg.IsWhitelisted(conn.RemoteAddr().String(), whitelist) {
		goLog.Debug(conn.RemoteAddr().String(), " not while list")
		return
	}

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
			realServerAddr, isok := realServerInfoMap[p.Id]
			if !isok {
				goLog.Error("not fount realservers ,id: ", p.Id)
				break
			}
			realServer, err := net.Dial("tcp", realServerAddr)
			if err != nil {
				goLog.Error("dial real server error", err)
				break
			}

			var c pipeprotocol.ClientInfo
			c.Conn = realServer
			c.Time = time.Now().Unix()

			func() {
				clientMapMutex.Lock()
				defer clientMapMutex.Unlock()
				clientConnections[p.Conn] = c
			}()

			go connectionRealServer(realServer, p.Conn, p.Id)

		}
		cInfo, _ := clientConnections[p.Conn]

		if p.CommandID == 100 {
			goLog.Info("close ", p.Id, " connecttion, commandid: ", p.CommandID)
			cInfo.Conn.Close()
			continue
		}

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
