package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/auth"
	httpCommon "github.com/thoratvinod/vi-chat/server/pkg/common/http"
	msgCommon "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/message"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebsocketHandler struct {
	MessageHandler message.MessageHandler
}

func (handler *WebsocketHandler) reader(conn *websocket.Conn, claims *auth.Token) {
	handler.MessageHandler.ConnManager.AddConnection(claims.UserID, conn)
	log.Printf("Established connection, userID=%+v", claims.UserID)
	for {
		msg := msgCommon.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			errResp := httpCommon.ErrorResponse{
				Message: fmt.Sprintf("invalid content received ; %s", err.Error()),
			}
			conn.WriteJSON(errResp)
			conn.Close()
			return
		}
		err = handler.MessageHandler.HandleMessage(&msg)
		if err != nil {
			errResp := httpCommon.ErrorResponse{
				Message: fmt.Sprintf("invalid content received ; %s", err.Error()),
			}
			conn.WriteJSON(errResp)
			continue
		}
	}
}

func (handler *WebsocketHandler) WSHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	jwtTokenStr, err := auth.ExtractJWTToken(r)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid auth JWT token provided ; %v", err.Error()), http.StatusForbidden)
		return
	}
	tokenClaims, err := auth.AuthenticateJWTToken(jwtTokenStr)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid auth JWT token provided ; %v", err.Error()), http.StatusForbidden)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("internal server error ; %v", err.Error()), http.StatusInternalServerError)
		return
	}
	handler.reader(ws, tokenClaims)
}
