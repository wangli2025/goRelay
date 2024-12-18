package pipetcpclient

import (
	"goRelay/pkg"
	"net"
	"sync"
)

var (
	goLog            pkg.Logger
	clientConnection net.Conn
	pipeClientConn   net.Conn
	err              error
	sendMutex        sync.Mutex
	sleepTimeSec     int = 3
)
