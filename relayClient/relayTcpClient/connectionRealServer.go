package relaytcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"io"
	"net"
	"time"
)

func connectionRealServer(realServerConn net.Conn, pConn string, id string) {
	goLog.Debug("connection real server ", realServerConn.RemoteAddr().String())

	realServerConn.SetReadDeadline(time.Now().Add(time.Hour * 12))
	realServerConn.SetWriteDeadline(time.Now().Add(time.Hour * 12))

	defer func() {
		clientMapMutex.Lock()
		defer clientMapMutex.Unlock()
		delete(clientConnections, pConn)
	}()

	for {
		buf := make([]byte, pipeprotocol.MaxPackageLen)

		n, err := realServerConn.Read(buf)
		if err != nil {
			if err != io.EOF {
				goLog.Error("from real server get data error", err)
			}
			func() {
				clientMapMutex.Lock()
				defer clientMapMutex.Unlock()
				goLog.Debug("delete client map", pConn)
				delete(clientConnections, pConn)
			}()
			break
		}
		msg := buf[:n]

		var p pipeprotocol.ClientProtocolInfo
		p.Id = id
		p.Buf = append(p.Buf, msg...)
		p.Conn = pConn

		jsonBuf, err := pkg.JsonMarshal(p)
		if err != nil {
			goLog.Error("json marshal error ", err)

			goLog.Info("delete client map ", pConn)
			delete(clientConnections, pConn)
			break
		}

		func() {
			clientMapMutex.Lock()
			defer clientMapMutex.Unlock()

			goLog.Debug("The request has been processed. Sending the payload back, payload size: ", len(jsonBuf))
			pipeprotocol.SendMessage(relayConn, jsonBuf)
		}()
	}
}
