package ws

// 批量发送
func (w *WsConnOb) BatchSend(raw []byte, ids []string) {

	for _, id := range ids {
		s := SendOb{
			Id:  id,
			Raw: raw,
		}
		w.outChan <- s
	}
}
