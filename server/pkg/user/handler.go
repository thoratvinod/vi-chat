package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	authCommon "github.com/thoratvinod/vi-chat/server/pkg/common/auth"
	httpCommon "github.com/thoratvinod/vi-chat/server/pkg/common/http"
)

type UserHandler struct {
	UserService *UserService
}

func (handler *UserHandler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
		return
	}

	err = handler.UserService.CreateUser(&user)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("internal server error ; %v", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (handler *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
		return
	}

	user, err := handler.UserService.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid login credentials provided ; %v", err.Error()), http.StatusBadRequest)
		return
	}

	jwtToken, err := authCommon.GenerateJWTToken(user.ID, user.Email)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("failed to generate JWT token ; %v", err.Error()), http.StatusBadRequest)
		return
	}

	resp := struct {
		JWTToken string `json:"jwtToken"`
	}{
		JWTToken: jwtToken,
	}

	json.NewEncoder(w).Encode(resp)
}
