package inithandlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/src/common"
	"github.com/thoratvinod/vi-chat/server/src/messagemgmt"
	"github.com/thoratvinod/vi-chat/server/src/usermgmt"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, claims *usermgmt.Token) {
	messagemgmt.ConnHandler.AddConnection(claims.UserID, conn)
	log.Printf("Established connection, userID=%+v", claims.UserID)
	for {
		msg := messagemgmt.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			errResp := common.NewErrorResponse(
				common.ErrInvalidInput,
				fmt.Sprintf("Invalid data received, err=%+v", err),
			)
			conn.WriteJSON(errResp)
			conn.Close()
			return
		}
		messagemgmt.ConnHandler.HandleMessage(msg)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	jwtTokenStr, err := common.ExtractJWTToken(r)
	if err != nil {
		errResp := common.NewErrorResponse(common.ErrAuthInvalidJWT, err.Error())
		common.WriteErrorJSON(w, errResp)
		return
	}

	token := usermgmt.Token{}
	tk, err := jwt.ParseWithClaims(
		jwtTokenStr,
		&token,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("Secret"), nil
		},
	)
	if err != nil {
		errResp := common.NewErrorResponse(common.ErrAuthInvalidJWT, err.Error())
		common.WriteErrorJSON(w, errResp)
		return
	}

	tokenClaims, ok := tk.Claims.(*usermgmt.Token)
	if !ok {
		errResp := common.NewErrorResponse(common.ErrAuthInvalidJWT, "")
		common.WriteErrorJSON(w, errResp)
		return
	}

	if !tk.Valid {
		errResp := common.NewErrorResponse(common.ErrAuthJWTExpired, "")
		common.WriteErrorJSON(w, errResp)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		errResp := common.NewErrorResponse(common.ErrInvalidInput, err.Error())
		common.WriteErrorJSON(w, errResp)
		return
	}
	reader(ws, tokenClaims)
}

func SetupRoutes(dbConn *gorm.DB) {

	//TODO: need to add mux router to handle 404
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong!")
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong!")
	})
	http.HandleFunc("/ws", wsEndpoint)
}
