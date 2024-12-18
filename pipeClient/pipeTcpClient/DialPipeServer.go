package pipetcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
	"time"
)

func ConnectToPipeServer(addr string) {
	for {
		pipeClientConn, err = net.Dial("tcp", addr)
		if err != nil {
			goLog.Error("Dial relay server error", err)
			time.Sleep(time.Second * time.Duration(sleepTimeSec))
			continue
		}
		goLog.Debug("to ", pipeClientConn.RemoteAddr().String(), " send relay Conn")
		pipeprotocol.SendMessage(pipeClientConn, []byte("relayConn"))

		pipeResponse := pipeprotocol.RecvMessgae(pipeClientConn)
		if nil == pipeResponse {
			goLog.Error("No available connections on the relay server.")
			continue
		}
		if string(pipeResponse) == "isok" {
			goLog.Info("Relay server connection established successfully.")
		}

		for {
			msg := pipeprotocol.RecvMessgae(pipeClientConn)
			if nil == msg {
				goLog.Error("Relay server connection is error")
				break
			}
			func() {
				sendMutex.Lock()
				defer sendMutex.Unlock()

				if nil == clientConnection {
					goLog.Error("relayClient Conn is null")
					return
				}

				var p pipeprotocol.ClientProtocolInfo
				err := pkg.JsonUnmarshal(msg, &p)
				if err != nil {
					goLog.Error("json unmarshal error", err)
				}
				pipeprotocol.SendMessage(clientConnection, msg)
			}()

		}
	}
}
