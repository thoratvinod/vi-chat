package connections

import (
	"fmt"
	"log"

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
		if !more {
			return
		}
		if msg.MsgType == message.GroupMessage {
			if err := cm.handleGroupMessage(msg); err != nil {
				log.Printf("got error while processing group message, err=%+v", err.Error())
			}
			continue
		}
		msgUserChan, err := cm.msgChannelMap.Get(msg.ReceiverID)
		if err != nil {
			log.Printf("error in startMessageProcessor, err=%v", err.Error())
			continue
		}
		msgUserChan <- msg
	}
}

func (cm *connManager) handleGroupMessage(msg *message.Message) error {
	if len(msg.MemberIDs) == 0 {
		return fmt.Errorf("member IDs array is empty")
	}
	errsIDs := []uint{}
	for _, memberID := range msg.MemberIDs {
		msgUserChan, err := cm.msgChannelMap.Get(memberID)
		if err != nil {
			errsIDs = append(errsIDs, memberID)
			continue
		}
		msgUserChan <- msg
	}
	if len(errsIDs) != 0 {
		return fmt.Errorf("error while sending message to userMsgChannel, memberIDs=%+v", errsIDs)
	}
	return nil
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
		conn, err := cm.connMap.Get(userID)
		if err != nil {
			log.Printf("error in processUserMessageChannel, err=%v", err.Error())
			continue
		}
		conn.WriteJSON(msg)
	}
}
