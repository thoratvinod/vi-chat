package connections

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
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
			userMsgChan, err := cm.msgChannelMap.Get(userID)
			if err != nil {
				log.Printf("error in connectionSanitizer, err=%+v", err.Error())
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
