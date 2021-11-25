package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

func (w *WsConnOb) writeLoop() {
	ticker := time.NewTicker(config.WS.PingPeriod)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case sendOb := <-w.outChan:
			if wc, ok := conn[sendOb.Id]; ok {
				sendLocal(wc, sendOb)
			}

		case <-w.closeChan:
			return
		case <-ticker.C:
			w.connect.SetWriteDeadline(time.Now().Add(writeWait))
			if err := w.connect.WriteMessage(websocket.PingMessage, nil); err != nil {
				w.close()
				return
			}
		}
	}
}

// 本地通讯
func sendLocal(w *WsConnOb, sendOb SendOb) {
	if err := w.connect.WriteMessage(websocket.TextMessage, sendOb.Raw); err != nil {
		return
	}
}
