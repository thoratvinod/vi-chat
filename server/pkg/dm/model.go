package dm

import (
	commonMsg "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"gorm.io/gorm"
)

type DirectMessage struct {
	gorm.Model
	SenderUserID   uint
	ReceiverUserID uint
	Content        string
	Status         commonMsg.MessageStatus
	TimeStamp      int64
}
