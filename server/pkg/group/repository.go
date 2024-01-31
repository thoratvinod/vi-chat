package group

import "gorm.io/gorm"

type GroupRepository struct {
	DB *gorm.DB
}

// func (repo *GroupRepository) CreateGroup(group *Group) (uint, error) {
// 	createdGroup := repo.DB.Create(group)
// 	if createdGroup.Error != nil {
// 		return 0, createdGroup.Error
// 	}
	
	
// }
