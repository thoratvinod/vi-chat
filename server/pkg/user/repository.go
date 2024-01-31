package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) CreateUser(user *User) error {
	var existingUser User
	res := repo.DB.Where("email = ?", user.Email).First(&existingUser)
	if res.Error == nil && res.RowsAffected > 0 {
		return fmt.Errorf("email already present, email=%+v", user.Email)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password encryption failed")
	}
	user.Password = string(encryptedPassword)

	createdUser := repo.DB.Save(user)
	if createdUser.Error != nil {
		return createdUser.Error
	}
	return nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*User, error) {
	user := User{}
	if err := repo.DB.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user with this email is not present")
	}
	return &user, nil
}
