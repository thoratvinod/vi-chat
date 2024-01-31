package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/auth"
	"github.com/thoratvinod/vi-chat/server/pkg/common/connmanager"
	httpCommon "github.com/thoratvinod/vi-chat/server/pkg/common/http"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connManager *connmanager.ConnManager

func init() {
	connManager = &connmanager.ConnManager{
		MsgQ:          make(chan *message.Message),
		MsgChannelMap: make(map[uint]chan *message.Message),
		ConnMap:       make(map[uint]*websocket.Conn),
	}
}

func reader(conn *websocket.Conn, claims *auth.Token) {
	connManager.AddConnection(claims.UserID, conn)
	log.Printf("Established connection, userID=%+v", claims.UserID)
	for {
		msg := message.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			errResp := httpCommon.ErrorResponse{
				Message: fmt.Sprintf("invalid content received ; %s", err.Error()),
			}
			conn.WriteJSON(errResp)
			conn.Close()
			return
		}
		connManager.HandleMessage(&msg)
	}
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
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
	reader(ws, tokenClaims)
}
