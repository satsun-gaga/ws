package ws

import "time"

var config Config

func init() {
	SetConfig(Config{
		WS: WsConfig{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			InChanLength:    500,
			OutChanLength:   500,
			MessageSize:     1024,
			PingPeriod:      10 * time.Second,

		},
	})
}

func SetConfig(c Config) {
	config = c
	initUpgrader()
}
