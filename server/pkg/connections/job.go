package connections

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

// removes inactive connection
func (cm *connManager) connectionSanitizer() {
	for {
		log.Println("Started connection sanitization")
		staleUserConns := []uint{}
		cm.connMap.Range(func(key, value any) bool {
			conn, ok := value.(*websocket.Conn)
			if !ok {
				log.Printf("type assertion failed in connectionSanitizer job for userID=%v", key)
				staleUserConns = append(staleUserConns, key.(uint))
				return true
			}
			if !isConnectionActive(conn) {
				staleUserConns = append(staleUserConns, key.(uint))
			}
			return true
		})
		for _, userID := range staleUserConns {
			cm.connMap.Delete(userID)
			val, ok := cm.msgChannelMap.Load(userID)
			if !ok {
				log.Printf("userMsgChan is not found for userID=%+v\n", userID)
				continue
			}
			userMsgChan, ok := val.(chan *message.Message)
			if !ok {
				log.Printf("type assertion failed for userMsgChan in connectionSanitizer job for userID=%+v\n", userID)
				continue
			}
			close(userMsgChan)
			log.Printf("deleted stale connection for userID=%v\n", userID)
			currActiveConns := cm.activeConns.Add(-1)
			if currActiveConns < 0 {
				log.Println("Error: active connection is less than zero")
			}
		}
		log.Println("Completed connection sanitization")
		time.Sleep(1 * time.Minute)
	}
}

func isConnectionActive(conn *websocket.Conn) bool {
	messageType := websocket.PingMessage
	payload := []byte("ping")
	err := conn.WriteMessage(messageType, payload)
	return err == nil
}
