package message

import (
	"fmt"

	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
	commonMsg "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/connections"
	"github.com/thoratvinod/vi-chat/server/pkg/dm"
	"github.com/thoratvinod/vi-chat/server/pkg/group"
)

type MessageHandler struct {
	DMService    *dm.DMService
	ConnManager  connections.ConnManager
	GroupService *group.GroupService
}

func (mp *MessageHandler) HandleMessage(msg *message.Message) error {
	switch msg.MsgType {
	case message.DirectMessage:
		return mp.handleDirectMessage(msg)
	case message.GroupMessage:
		return mp.handleGroupMessage(msg)
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
		Status:         commonMsg.MsgStatusCreated,
	})
	if err != nil {
		return err
	}
	// notify message
	mp.ConnManager.NotifyMsg(msg)
	return nil
}

func (mp *MessageHandler) handleGroupMessage(msg *message.Message) error {
	// save message in DB
	err := mp.GroupService.CreateGroupMessage(&group.GroupMessage{
		SendUserID: msg.SenderID,
		GroupID:    msg.ReceiverID,
		Content:    msg.Content,
		TimeStamp:  msg.Timestamp,
	})
	if err != nil {
		return err
	}

	users, err := mp.GroupService.GetGroupMembers(msg.ReceiverID)
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.ID == msg.SenderID {
			continue
		}
		msg.MemberIDs = append(msg.MemberIDs, user.ID)
	}
	// notify message
	mp.ConnManager.NotifyMsg(msg)
	return nil
}
