package v1

import (
	"net/http"

	"github.com/thoratvinod/vi-chat/server/pkg/websocket"
)

func SetWebsocketRoute() {
	http.HandleFunc("/ws", websocket.WebsocketHandler)
}
