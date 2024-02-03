package v1

import (
	"net/http"

	"github.com/thoratvinod/vi-chat/server/pkg/connections"
	"github.com/thoratvinod/vi-chat/server/pkg/dm"
	"github.com/thoratvinod/vi-chat/server/pkg/message"
	"github.com/thoratvinod/vi-chat/server/pkg/websocket"
	"gorm.io/gorm"
)

func SetWebsocketRoute(dbConn *gorm.DB) {

	msgHandler := message.MessageHandler{
		ConnManager: connections.NewConnManager(1000),
		DMService: &dm.DMService{
			Repo: &dm.DMRepository{
				DB: dbConn,
			},
		},
	}
	websocketHandler := websocket.WebsocketHandler{
		MessageHandler: msgHandler,
	}
	http.HandleFunc("/ws", websocketHandler.WSHandler)
}
