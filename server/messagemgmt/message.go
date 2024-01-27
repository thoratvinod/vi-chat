package messagemgmt

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	MsgType    MessageType `json:"msgType"`
	SenderID   uint        `json:"senderID"`
	ReceiverID uint        `json:"receiverID"`
	// Timestamp  time.Time `timeStamp`
	Content string `json:"content"`
	// it can be id of group or some other entity
	// ResourseID uint `json:"resourceID"`
}

var ConnHandler *ConnectionHandler

func init() {
	ConnHandler = NewConnectionHandler()
	go ConnHandler.StartMessageProcessor()
}

type ConnectionHandler struct {
	// this message queue is responsible for consuming message and again distribute to specifc user channels
	msgQ chan Message
	// userID to channel of message
	msgChannelMap map[uint]chan Message
	// userID to connection map
	connMap map[uint]*websocket.Conn
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		msgQ:          make(chan Message, 1000),
		msgChannelMap: map[uint]chan Message{},
		connMap:       map[uint]*websocket.Conn{},
	}
}

func (ch *ConnectionHandler) AddConnection(userID uint, conn *websocket.Conn) {
	ch.msgChannelMap[userID] = make(chan Message, 10)
	ch.connMap[userID] = conn
	go ch.processUserMessageChannel(userID, ch.msgChannelMap[userID])
}

func (ch *ConnectionHandler) HandleMessage(msg Message) {
	log.Printf("handling message senderID=%v receiverID=%v", msg.SenderID, msg.ReceiverID)
	ch.msgQ <- msg
}

func (ch *ConnectionHandler) StartMessageProcessor() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic occurred, starting StartMessageProcessor goroutinue again")
			go ch.StartMessageProcessor()
		}
	}()
	for {
		msg, more := <-ch.msgQ
		ch.msgChannelMap[msg.ReceiverID] <- msg
		if !more {
			break
		}
	}
}

func (ch *ConnectionHandler) processUserMessageChannel(userID uint, userMsgQ chan Message) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic occurred, starting processUserMessageChannel goroutinue again, userID=%+v\n", userID)
			go ch.processUserMessageChannel(userID, userMsgQ)
		}
	}()
	for {
		msg, more := <-userMsgQ
		log.Printf("received message for userID=%v, message=%v", userID, msg)
		conn, ok := ch.connMap[userID]
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
