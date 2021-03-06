package ws

import "github.com/gorilla/websocket"

// 读取队列
func (w *WsConnOb) readLoop() {

	w.connect.SetReadLimit(config.WS.MessageSize)
	// w.connect.SetReadDeadline(time.Now().Add(config.WS.PongWait))
	for {
		// 读一个message
		_, data, err := w.connect.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			w.close()
			return
		}

		select {
		case w.inChan <- data:
		case <-w.closeChan:
			return
		}
	}
}
