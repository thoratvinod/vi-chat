package connections

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

func (cm *connManager) startMessageProcessor() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occurred, starting StartMessageProcessor goroutinue again")
			go cm.startMessageProcessor()
		}
	}()
	for {
		msg, more := <-cm.msgQ
		val, ok := cm.msgChannelMap.Load(msg.ReceiverID)
		if !ok {
			log.Printf("can't find user's message channel in MsgChannelMap, userID=%+v, msg=%+v\n",
				msg.ReceiverID, msg)
			continue
		}
		msgUserChan, ok := val.(chan *message.Message)
		if !ok {
			log.Printf("type assertion failed in StartMessageProcessor, msg=%+v\n", msg)
			continue
		}
		msgUserChan <- msg
		if !more {
			break
		}
	}
}

func (cm *connManager) processUserMessageChannel(userID uint, userMsgQ chan *message.Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic occurred, starting processUserMessageChannel goroutinue again, userID=%+v\n", userID)
			go cm.processUserMessageChannel(userID, userMsgQ)
		}
	}()
	for {
		msg, more := <-userMsgQ
		if !more {
			log.Printf("stopping processUserMessageChannel goroutine for userID=%v\n", userID)
			return
		}
		log.Printf("received message for userID=%v, message=%v", userID, msg)
		val, ok := cm.connMap.Load(userID)
		if !ok {
			log.Printf("can't find connection for userID, userID=%+v, msg=%+v\n", userID, msg)
			return
		}
		conn, ok := val.(*websocket.Conn)
		if !ok {
			log.Printf("type assertion failed in processUserMessageChannel, msg=%+v\n", msg)
			return
		}
		conn.WriteJSON(msg)
	}
}
