package msg

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

type ConnManager struct {
	// this message queue is responsible for consuming message and again distribute to specifc user channels
	MsgQ chan *message.Message
	// userID to channel of message
	MsgChannelMap map[uint]chan *message.Message
	// userID to connection map
	ConnMap map[uint]*websocket.Conn
}

func (cm *ConnManager) AddConnection(userID uint, conn *websocket.Conn) {
	cm.MsgChannelMap[userID] = make(chan *message.Message, 10)
	cm.ConnMap[userID] = conn
	go cm.processUserMessageChannel(userID, cm.MsgChannelMap[userID])
}

func (cm *ConnManager) HandleMessage(msg *message.Message) {
	log.Printf("handling message senderID=%v receiverID=%v", msg.SenderID, msg.ReceiverID)
	cm.MsgQ <- msg
}

func (cm *ConnManager) StartMessageProcessor() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occurred, starting StartMessageProcessor goroutinue again")
			go cm.StartMessageProcessor()
		}
	}()
	for {
		msg, more := <-cm.MsgQ
		cm.MsgChannelMap[msg.ReceiverID] <- msg
		if !more {
			break
		}
	}
}

func (cm *ConnManager) processUserMessageChannel(userID uint, userMsgQ chan *message.Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic occurred, starting processUserMessageChannel goroutinue again, userID=%+v\n", userID)
			go cm.processUserMessageChannel(userID, userMsgQ)
		}
	}()
	for {
		msg, more := <-userMsgQ
		log.Printf("received message for userID=%v, message=%v", userID, msg)
		conn, ok := cm.ConnMap[userID]
		if !ok {
			log.Println("Connection is not present for now")
			continue
		}
		conn.WriteJSON(msg)
		if !more {
			break
		}
	}
}
