package relaytcpserver

import (
	"time"
)

func DeleteHistoryClientMessage() {
	for {
		goLog.Debug("Start performing historical client message cleanup.")
		time.Sleep(time.Minute * time.Duration(sleepDeleteHistoryClientMinute))
		func() {
			clientMapLock.Lock()
			defer clientMapLock.Unlock()
			for k, v := range clientMap {
				goLog.Debug("find clients: ", k, v.Time)
				if (v.Time + (historyClientTimeout)) < time.Now().Unix() {
					goLog.Debug("close ", k, " infos")
					v.Conn.Close()
				}
			}
		}()
	}
}
