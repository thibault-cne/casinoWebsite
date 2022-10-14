package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsUpgrader(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	conn, err := wsUpgrader.Upgrade(writer, request, nil)

	return conn, err
}
