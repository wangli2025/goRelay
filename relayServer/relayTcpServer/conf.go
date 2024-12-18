package relaytcpserver

import (
	pipeprotocol "goRelay/pipeProtocol"
	"goRelay/pkg"
	"net"
	"sync"
)

var (
	pipeClient                     net.Conn
	err                            error
	clientMap                      map[string]pipeprotocol.ClientInfo
	pipeLock                       sync.Mutex
	clientMapLock                  sync.Mutex
	goLog                          pkg.Logger
	sleepDeleteHistoryClientMinute int   = 1440
	historyClientTimeout           int64 = (3 * 1440 * 60)
	sleepTimeSecond                int   = 3
)
