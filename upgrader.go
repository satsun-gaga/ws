package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type wu struct {
	websocket.Upgrader
}

var Upgrader wu

func initUpgrader() {
	Upgrader = wu{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  config.WS.ReadBufferSize,
			WriteBufferSize: config.WS.WriteBufferSize,
		},
	}
}

// 升级websocket，生成id
func (u wu) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*WsConnOb, error) {

	socket, err := u.Upgrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	wsId := getRandomString(8)
	connOb := WsConnOb{
		Id:        wsId,
		connect:   socket,
		inChan:    make(chan []byte, config.WS.InChanLength),
		outChan:   make(chan SendOb, config.WS.OutChanLength),
		closeChan: make(chan struct{}),
	}

	go connOb.writeLoop()
	go connOb.readLoop()

	conn[wsId] = &connOb
	return &connOb, err
}

