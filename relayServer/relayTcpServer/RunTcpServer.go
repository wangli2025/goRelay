package relaytcpserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"io"
	"net"
	"time"
)

func RunTcpServer(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		goLog.Error("listen ", addr, " error ", err)
		return
	}
	goLog.Info("listen ", addr, " successful")

	for {
		conn, err := listen.Accept()
		if err != nil {
			goLog.Error("accept error", err)
			continue
		}
		go worker(conn)
	}
}

func worker(conn net.Conn) {
	goLog.Debug(conn.RemoteAddr().String(), " is connection")

	func() {
		clientMapLock.Lock()
		defer clientMapLock.Unlock()
		var cInfo pipeprotocol.ClientInfo
		cInfo.Conn = conn
		cInfo.Time = time.Now().Unix()
		clientMap[conn.RemoteAddr().String()] = cInfo
	}()

	for {
		buf := make([]byte, pipeprotocol.MaxPackageLen)
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				goLog.Error("conn read relay server error,error: ", err)
			}

			func() {
				clientMapLock.Lock()
				defer clientMapLock.Unlock()
				delete(clientMap, conn.RemoteAddr().String())
			}()
			break
		}

		msg := buf[:n]
		var p pipeprotocol.ClientProtocolInfo
		p.Conn = conn.RemoteAddr().String()
		p.Buf = append(p.Buf, msg...)

		jsonBuf, err := pkg.JsonMarshal(p)
		if err != nil {
			goLog.Error("json marshal error", err)
			return
		}

		func() {
			pipeLock.Lock()
			defer pipeLock.Unlock()
			if nil == pipeClient {
				goLog.Error("pipeClientConn is nil")
				return
			}
			goLog.Debug(p.Conn, " The message has been sent to the pipeline.")
			pipeprotocol.SendMessage(pipeClient, jsonBuf)
		}()

		cInfo := clientMap[conn.RemoteAddr().String()]
		cInfo.Time = time.Now().Unix()

		func() {
			clientMapLock.Lock()
			defer clientMapLock.Unlock()
			clientMap[conn.RemoteAddr().String()] = cInfo
		}()
	}
}
