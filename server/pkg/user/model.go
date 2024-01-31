package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password"`
}
