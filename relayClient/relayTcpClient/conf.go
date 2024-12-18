package relaytcpclient

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
	"sync"
)

var (
	clientConnections   map[string]pipeprotocol.ClientInfo
	goLog               pkg.Logger
	clientMapMutex      sync.Mutex
	relayConn           net.Conn
	clientCheckInterval int   = 1440
	clientTimeout       int64 = (3 * 1440 * 60)
)
