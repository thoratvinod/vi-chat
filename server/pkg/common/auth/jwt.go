package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwtSecret = []byte("Secret")
)

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

func GenerateJWTToken(userID uint, email string) (string, error) {
	expiresAt := time.Now().Add(1000 * time.Minute).Unix()
	tk := Token{
		UserID: userID,
		Email:  email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &tk)

	jwtStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token ; %v", err.Error())
	}
	return jwtStr, nil
}

func AuthenticateJWTToken(jwtStr string) (*Token, error) {
	token := Token{}
	tk, err := jwt.ParseWithClaims(
		jwtStr,
		&token,
		func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error while parsing JWT token ; %v", err.Error())
	}

	tokenClaims, ok := tk.Claims.(*Token)
	if !ok {
		return nil, fmt.Errorf("invalid claims received ; type assertion failed")
	}
	if !tk.Valid {
		return nil, fmt.Errorf("invalid JWT token provided")
	}
	return tokenClaims, nil
}

func ExtractJWTToken(r *http.Request) (string, error) {
	// get auth jwt token
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return "", fmt.Errorf("authorization JWT token not provided")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 || len(splitToken[1]) == 0 {
		return "", fmt.Errorf("invalid JWT token received")
	}
	return splitToken[1], nil
}
