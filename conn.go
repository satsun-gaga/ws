package ws

// 连接信息
var conn map[string]*WsConnOb


func init() {
	conn = make(map[string]*WsConnOb)
}
