package group

import (
	commonMsg "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/user"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name            string       `json:"name" gorm:"uniqueIndex"`
	Desc            string       `json:"description"`
	Members         []*user.User `gorm:"many2many:group_members;"`
	CreatedBy       *user.User   `gorm:"foreignKey:CreatedByUserID"`
	MembersUserIDs  []uint       `gorm:"-" json:"membersUserIDs"`
	CreatedByUserID uint         `grom:"-" json:"createdByUserID"`
}

type GroupMessage struct {
	gorm.Model
	SendUserID uint
	GroupID    uint
	Content    string
	Statuses   []*GroupMessageStatus `gorm:"foreignKey:GroupMessageID"`
	TimeStamp  int64
}

type GroupMessageStatus struct {
	gorm.Model
	GroupMessageID uint
	GroupMessage   GroupMessage `gorm:"foreignKey:GroupMessageID"`
	UserID         uint
	Status         commonMsg.MessageStatus
}
