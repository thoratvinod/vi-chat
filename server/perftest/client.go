package perftest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type ClientConnectionHandler struct {
	JWTTokenMap map[uint]string
	ConnMap     map[uint]*websocket.Conn
}

func NewClientConnectionHandler(JWTTokenMap map[uint]string) *ClientConnectionHandler {
	return &ClientConnectionHandler{
		JWTTokenMap: JWTTokenMap,
		ConnMap:     map[uint]*websocket.Conn{},
	}
}

func (cch *ClientConnectionHandler) EstablishConnections(baseURL string) {

	for userID, token := range cch.JWTTokenMap {
		headers := http.Header{}
		headers.Add("Authorization", fmt.Sprintf("Bearer %v", token))
		conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%v/ws", baseURL), headers)
		if err != nil {
			log.Printf("connection establishment failed, userID=%v, err=%+v\n", userID, err.Error())
			continue
		}
		cch.ConnMap[userID] = conn
	}
}
