package relaytcpserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
	"time"
)

func ConnectPipeServer(pipeServerAddr string) {
	for {
		pipeClient, err = net.Dial("tcp", pipeServerAddr)
		if err != nil {
			goLog.Error("Unable to connect to the PipeServer server. ", err)
			pipeClient = nil
			time.Sleep(time.Duration(sleepTimeSecond) * time.Second)
			continue
		}
		goLog.Info("dial ", pipeServerAddr, " successful")

		relayWorker(pipeClient)
	}
}

func relayWorker(pipeClient net.Conn) {
	for {
		msg := pipeprotocol.RecvMessgae(pipeClient)
		if nil == msg {
			goLog.Error("recv mesage error,stop relay tcp connection.")
			break
		}

		var p pipeprotocol.ClientProtocolInfo
		pkg.JsonUnmarshal(msg, &p)
		goLog.Debug(p.Conn, " The message has returned.")

		cInfo, isok := clientMap[p.Conn]
		if !isok {
			goLog.Error("not fount client conn,error,conn: ", p.Conn)
			continue
		}

		sendLen := 0
		for sendLen < len(p.Buf) {
			n, err := cInfo.Conn.Write(p.Buf[sendLen:])
			if err != nil {
				goLog.Error("write real client error ", err)
				break
			}
			sendLen += n
		}
	}
}
