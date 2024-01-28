package usermgmt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userManagement struct {
	db *gorm.DB
}

func (um *userManagement) CreateUser(user *User) error {

	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email or password cannot be empty, email=%+v", user.Email)
	}

	var existingUser User
	res := um.db.Where("email = ?", user.Email).First(&existingUser)
	if res.Error == nil && res.RowsAffected > 0 {
		return fmt.Errorf("email already present, email=%+v", user.Email)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password encryption failed")
	}
	user.Password = string(encryptedPassword)

	createdUser := um.db.Save(user)
	if createdUser.Error != nil {
		return createdUser.Error
	}
	return nil
}

func (um *userManagement) FindOne(email string, password string) (*User, error) {

	user := User{}
	if err := um.db.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user with this email is not present")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, fmt.Errorf("invalid login credentials. Please try again")
	}

	return &user, nil
}
