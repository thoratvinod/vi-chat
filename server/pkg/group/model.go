package group

import (
	"github.com/thoratvinod/vi-chat/server/pkg/user"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name            string       `json:"name" gorm:"uniqueIndex"`
	Desc            string       `json:"description"`
	Members         []*user.User `gorm:"many2many:group_members;"`
	CreatedBy       *user.User   `gorm:"foreignKey:CreatedByUserID"`
	MembersUserID   []uint       `gorm:"-" json:"membersUserID"`
	CreatedByUserID uint
}
