package usermgmt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string
}

type Token struct {
	UserID         uint
	Email          string
	StandardClaims *jwt.StandardClaims
}

func (t *Token) Valid() error {
	expirationTime := time.Unix(t.StandardClaims.ExpiresAt, 0)
	if time.Now().After(expirationTime) {
		return fmt.Errorf("JWT token is expired")
	}
	return nil
}
