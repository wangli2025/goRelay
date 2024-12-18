package pipeserver

import (
	"goRelay/pkg"
	"net"
	"sync"
)

var relayConn net.Conn = nil
var clientConn net.Conn = nil
var sendLock sync.Mutex
var goLog pkg.Logger
