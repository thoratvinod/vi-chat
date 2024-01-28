package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)


type chatManager struct {
	// userid -> connection
	conns map[string]*websocket.Conn
}

func NewChatManager() *chatManager {
	return &chatManager{
		conns: make(map[string]*websocket.Conn),
	}
}

func (cm * chatManager) Send(message string, userID string) error {
	
	conn, ok := cm.conns[userID]
	if !ok {
		return fmt.Errorf("connection for userID is not present, userId=%v", userID)
	}

	conn.WriteMessage(websocket.TextMessage, []byte(message))
	return nil
}

func (cm *chatManager) HandleReceivedMessage(message string, userID string) {
	
}



