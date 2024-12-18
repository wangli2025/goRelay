package pipeprotocol

import (
	"encoding/binary"
	"net"
)

type ClientProtocolInfo struct {
	Conn string `json:conn1`
	Buf  []byte `json:"buf"`
}

type ClientInfo struct {
	Conn net.Conn
	Time int64
}

func SendMessage(conn net.Conn, message []byte) {

	msg := Encode(message)
	msgLen := uint32(len(msg))

	headBuf := make([]byte, headerLen)
	binary.LittleEndian.PutUint32(headBuf, msgLen)

	sendLen := 0
	for sendLen < headerLen {
		n, err := conn.Write(headBuf[sendLen:])
		if err != nil {
			return
		}
		sendLen += n
	}

	sendLen = 0
	for sendLen < int(msgLen) {
		n, err := conn.Write([]byte(msg)[sendLen:])
		if err != nil {
			return
		}
		sendLen += n
	}
}

func RecvMessgae(conn net.Conn) []byte {

	var msgLen uint32
	headBuf := make([]byte, headerLen)

	recvLen := 0
	for recvLen < headerLen {
		n, err := conn.Read(headBuf[recvLen:])
		if err != nil {
			return nil
		}
		recvLen += n
	}

	msgLen = binary.LittleEndian.Uint32(headBuf)

	buf := make([]byte, msgLen)

	recvLen = 0
	for recvLen < int(msgLen) {
		n, err := conn.Read(buf[recvLen:])
		if err != nil {
			return nil
		}
		recvLen += n
	}

	deBuf := Decode(buf[:])

	return deBuf
}
