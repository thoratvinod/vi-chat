package group

import (
	commonMsg "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/user"
	"gorm.io/gorm"
)

type GroupRepository struct {
	DB *gorm.DB
}

func (repo *GroupRepository) SaveGroup(group *Group) (uint, error) {
	createdGroup := repo.DB.Save(group)
	if createdGroup.Error != nil {
		return 0, createdGroup.Error
	}
	return group.ID, nil
}

func (repo *GroupRepository) GetGroupMembers(groupID uint) ([]*user.User, error) {
	var group Group
	if err := repo.DB.Preload("Members").First(&group, groupID).Error; err != nil {
		return nil, err
	}
	return group.Members, nil
}

func (repo *GroupRepository) AddMembers(groupID uint, memberIDs []uint) error {
	var group Group
	if err := repo.DB.First(&group, groupID).Error; err != nil {
		return err
	}
	var members []*user.User
	if err := repo.DB.Where(memberIDs).Find(&members).Error; err != nil {
		return err
	}
	return repo.DB.Model(&group).Association("Members").Append(members)
}

func (repo *GroupRepository) RemoveMembers(groupID uint, memberIDs []uint) error {
	var group Group
	if err := repo.DB.First(&group, groupID).Error; err != nil {
		return err
	}
	var members []*user.User
	if err := repo.DB.Where(memberIDs).Find(&members).Error; err != nil {
		return err
	}
	return repo.DB.Model(&group).Association("Members").Delete(members)
}

func (repo *GroupRepository) CreateGroupMessage(grpMsg *GroupMessage) error {
	return repo.DB.Create(grpMsg).Error
}

func (repo *GroupRepository) UpdateGroupMessageStatus(grpMsgID uint, status commonMsg.MessageStatus) error {
	return repo.DB.Model(&GroupMessage{}).Where("ID = ?", grpMsgID).Update("status", status).Error
}
