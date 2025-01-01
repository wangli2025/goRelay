package pipeserver

import (
	"goRelay/pkg"
	"net"
	"sync"
)

var relayConn net.Conn = nil
var clientConnMap map[string]net.Conn
var clientConnMapLock sync.Mutex
var sendLock sync.Mutex
var goLog *pkg.Logger
