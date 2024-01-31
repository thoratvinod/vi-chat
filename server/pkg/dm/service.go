package dm

import "fmt"

type DMService struct {
	Repo *DMRepository
}

func (service *DMService) SaveMessage(dm *DirectMessage) error {
	if dm.SenderUserID == 0 || dm.ReceiverUserID == 0 || dm.Content == "" {
		return fmt.Errorf("invalid DM provided")
	}
	return service.Repo.SaveDirectMessage(dm)
}
