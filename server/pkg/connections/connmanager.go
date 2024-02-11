package connections

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

type ConnManager interface {
	// Add connection
	AddConnection(userID uint, conn *websocket.Conn)
	// Notifies msg to that user
	// If user is inactive it will log error for that
	NotifyMsg(msg *message.Message)
}

type connManager struct {
	// Queue is responsible for consuming message and again distribute to specifc user channels
	msgQ chan *message.Message
	// UserID to user specific channel mapping
	msgChannelMap MessageChannelMap
	// UserID to websocket connection mapping
	connMap ConnectionMap
	// MaxLimit for connections
	maxLimit int32
	// active connections
	activeConns atomic.Int32
}

func NewConnManager(maxLimit int32) ConnManager {
	connMgr := connManager{
		msgQ:          make(chan *message.Message),
		msgChannelMap: MessageChannelMap{syncMap: sync.Map{}},
		connMap:       ConnectionMap{syncMap: sync.Map{}},
		maxLimit:      maxLimit,
	}
	// message processor
	go connMgr.startMessageProcessor()
	// sanitizer job
	go connMgr.connectionSanitizer()
	return &connMgr
}

func (cm *connManager) AddConnection(userID uint, conn *websocket.Conn) {
	userMsgChan := make(chan *message.Message, 10)
	cm.msgChannelMap.Put(userID, userMsgChan)
	cm.connMap.Put(userID, conn)
	go cm.processUserMessageChannel(userID, userMsgChan)
	cm.activeConns.Add(1)
}

func (cm *connManager) NotifyMsg(msg *message.Message) {
	log.Printf("handling message senderID=%v receiverID=%v", msg.SenderID, msg.ReceiverID)
	cm.msgQ <- msg
}
