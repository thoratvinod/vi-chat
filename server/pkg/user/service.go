package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *UserRepository
}

func (service *UserService) AuthenticateUser(email, password string) (*User, error) {
	user, err := service.UserRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(user *User) error {
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email or password cannot be empty")
	}
	return service.UserRepo.CreateUser(user)
}
