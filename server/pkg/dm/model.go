package dm

import "gorm.io/gorm"

type DMStatus uint

const (
	DMStatusNone DMStatus = iota
	DMStatusCreated
	DMStatusSent
	DMStatusSeen
)

type DirectMessage struct {
	gorm.Model
	SenderUserID   uint
	ReceiverUserID uint
	Content        string
	Status         DMStatus
	TimeStamp      int64
}
