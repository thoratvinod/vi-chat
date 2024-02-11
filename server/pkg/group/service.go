package group

import (
	"fmt"

	commonMsg "github.com/thoratvinod/vi-chat/server/pkg/common/message"
	"github.com/thoratvinod/vi-chat/server/pkg/user"
)

type GroupService struct {
	Repo *GroupRepository
}

func (service *GroupService) SaveGroup(group *Group) error {
	forCreate := false
	if group.ID == 0 {
		forCreate = true
	}
	_, err := service.Repo.SaveGroup(group)
	if err != nil {
		return fmt.Errorf("SaveGroup failed ; %+v", err.Error())
	}
	if forCreate {
		return service.AddMembers(group.ID, group.MembersUserIDs)
	}
	return nil
}

func (service *GroupService) GetGroupMembers(groupID uint) ([]*user.User, error) {
	return service.Repo.GetGroupMembers(groupID)
}

func (service *GroupService) AddMembers(groupID uint, memberIDs []uint) error {
	if err := service.Repo.AddMembers(groupID, memberIDs); err != nil {
		return fmt.Errorf("add member failed ; %+v", err.Error())
	}
	return nil
}

func (service *GroupService) RemoveMembers(groupID uint, memberIDs []uint) error {
	if err := service.Repo.RemoveMembers(groupID, memberIDs); err != nil {
		return fmt.Errorf("remove member failed ; %+v", err.Error())
	}
	return nil
}

func (service *GroupService) CreateGroupMessage(grpMsg *GroupMessage) error {
	if err := service.Repo.CreateGroupMessage(grpMsg); err != nil {
		return fmt.Errorf("CreateGroupMessage failed ; %v", err.Error())
	}
	return nil
}

func (service *GroupService) UpdateGroupMessageStatus(grpMsgID uint, status commonMsg.MessageStatus) error {
	if err := service.Repo.UpdateGroupMessageStatus(grpMsgID, status); err != nil {
		return fmt.Errorf("UpdateGroupMessageStatus failed ; %v", err.Error())
	}
	return nil
}
