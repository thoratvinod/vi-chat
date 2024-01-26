package usermgmt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/thoratvinod/vi-chat/common"
	"gorm.io/gorm"
)

var um *userManagement

const (
	SignupPath = "/user/signup"
	LoginPath  = "/user/login"
)

func SetupRoutes(inputdb *gorm.DB) {
	um = &userManagement{
		db: inputdb,
	}
	common.RegisterRoute(http.MethodPost, SignupPath, signup)
	common.RegisterRoute(http.MethodPost, LoginPath, login)
}

func signup(w http.ResponseWriter, r *http.Request) (any, common.ErrorResponse) {
	req := SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, common.NewErrorResponse(
			common.ErrSignupFailed,
			fmt.Sprintf("Invalid request received; controller=%s err=%s", SignupPath, err.Error()),
		)
	}
	err = um.CreateUser(&User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		return nil, common.NewErrorResponse(
			common.ErrSignupFailed,
			fmt.Sprintf("Failed to create user, controller=%s, err=%s", SignupPath, err.Error()),
		)
	}
	return &SignupResponse{}, nil
}

func login(w http.ResponseWriter, r *http.Request) (any, common.ErrorResponse) {
	req := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, common.NewErrorResponse(
			common.ErrLoginFailed,
			fmt.Sprintf("Invalid request received; controller=%s err=%s", LoginPath, err.Error()),
		)
	}
	user, err := um.FindOne(req.Email, req.Password)
	if err != nil {
		return nil, common.NewErrorResponse(
			common.ErrLoginFailed,
			fmt.Sprintf("Failed to login user, controller=%s, err=%s", LoginPath, err.Error()),
		)
	}

	expiresAt := time.Now().Add(1000 * time.Minute).Unix()

	// form jwt token
	tk := &token{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return nil, common.NewErrorResponse(
			common.ErrLoginFailed,
			fmt.Sprintf("Failed to generate signed jwt token, controller=%s, err=%s", LoginPath, err.Error()),
		)
	}

	resp := LoginResponse{
		JWTToken: tokenString,
	}
	return &resp, nil
}
