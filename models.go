package ws

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

const (
	writeWait = 10 * time.Second // 允许等待的写入时间
)

// 系统配置
type Config struct {
	WS    WsConfig    // ws配置
}

type WsConfig struct {

	ReadBufferSize  int
	WriteBufferSize int

	InChanLength    int
	OutChanLength   int
	MessageSize     int64
	PingPeriod      time.Duration

}


// 发送数据对象
type SendOb struct {
	Id  string
	Raw []byte
}

type sendResponse struct {
	Status bool   `json:"status"`
	Msg    string `json:"msg"`
}

// ws连接对象
type WsConnOb struct {
	Id        string
	connect   *websocket.Conn
	inChan    chan []byte
	outChan   chan SendOb
	closeChan chan struct{}
	mutex     sync.Mutex
	isClosed  bool
}
