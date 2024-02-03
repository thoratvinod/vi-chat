package message

import (
	"fmt"

	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/connections"
	"github.com/thoratvinod/vi-chat/server/pkg/dm"
)

type MessageHandler struct {
	DMService   *dm.DMService
	ConnManager connections.ConnManager
}

func (mp *MessageHandler) HandleMessage(msg *message.Message) error {
	switch msg.MsgType {
	case message.DirectMessage:
		return mp.handleDirectMessage(msg)
	default:
		return fmt.Errorf("invalid message type provide, msgType=%v", msg.MsgType)
	}
}

func (mp *MessageHandler) handleDirectMessage(msg *message.Message) error {
	// save message in DB
	err := mp.DMService.SaveMessage(&dm.DirectMessage{
		SenderUserID:   msg.SenderID,
		ReceiverUserID: msg.ReceiverID,
		TimeStamp:      msg.Timestamp,
		Content:        msg.Content,
		Status:         dm.DMStatusCreated,
	})
	if err != nil {
		return err
	}
	// notify message
	mp.ConnManager.NotifyMsg(msg)
	return nil
}
