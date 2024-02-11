package v1

import (
	"net/http"

	"github.com/thoratvinod/vi-chat/server/pkg/group"
)

func SetupGroupRoutes(handler *group.GroupHandler) {
	http.HandleFunc("/group", handler.SaveGroup)
	http.HandleFunc("/group/member", handler.AddMembers)
}
