package relaytcpclient

import (
	"time"
)

func DeleteHistoryClientMessage() {
	for {
		goLog.Debug("Start performing historical client message cleanup.")
		time.Sleep(time.Minute * time.Duration(clientCheckInterval))
		func() {
			clientMapMutex.Lock()
			defer clientMapMutex.Unlock()
			for k, v := range clientConnections {
				goLog.Debug("find clients: ", k, v.Time)
				if (v.Time + (clientTimeout)) < time.Now().Unix() {
					goLog.Debug("close ", k, " infos")
					v.Conn.Close()
				}
			}
		}()
	}
}
