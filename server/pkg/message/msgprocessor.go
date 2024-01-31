package msg

import (
	"fmt"

	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/dm"
)

type MsgProcessor struct {
	dmService dm.DMService
}

func (mp *MsgProcessor) Start() {

}

func (mp *MsgProcessor) HandleMessage(msg *message.Message) error {
	switch msg.MsgType {
	case message.DirectMessage:
		return mp.handleDirectMessage(msg)
	default:
		return fmt.Errorf("invalid message type provide, msgType=%v", msg.MsgType)
	}
}

func (mp *MsgProcessor) handleDirectMessage(msg *message.Message) error {
	err := mp.dmService.SaveMessage(&dm.DirectMessage{
		SenderUserID:   msg.SenderID,
		ReceiverUserID: msg.ReceiverID,
		TimeStamp:      msg.Timestamp,
		Content:        msg.Content,
		Status:         dm.DMStatusCreated,
	})
	if err != nil {
		return err
	}
	return nil
}
