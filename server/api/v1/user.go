package v1

import (
	"net/http"

	"github.com/thoratvinod/vi-chat/server/pkg/user"
)

func SetupUserRoutes(handler *user.UserHandler) {
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/login", handler.LoginHandler)
}
