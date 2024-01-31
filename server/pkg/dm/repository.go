package dm

import (
	"fmt"

	"gorm.io/gorm"
)

type DMRepository struct {
	DB *gorm.DB
}

func (repo *DMRepository) SaveDirectMessage(dm *DirectMessage) error {
	createdDM := repo.DB.Save(dm)
	if createdDM.Error != nil {
		return fmt.Errorf("failed to save direct message ; %+v", createdDM.Error.Error())
	}
	return nil
}
