package main

import (
	"log"
	"net/http"

	api "github.com/thoratvinod/vi-chat/server/api/v1"
	"github.com/thoratvinod/vi-chat/server/inithandlers"
	"github.com/thoratvinod/vi-chat/server/pkg/group"
	"github.com/thoratvinod/vi-chat/server/pkg/user"
)

func main() {
	dbConn, err := inithandlers.InitDatabase()
	if err != nil {
		log.Fatalf("DB connection attempt failed, err=%+v", err)
	}
	api.SetupPingRoute()
	api.SetWebsocketRoute(dbConn)
	userHandler := user.UserHandler{
		UserService: &user.UserService{
			UserRepo: &user.UserRepository{
				DB: dbConn,
			},
		},
	}
	api.SetupUserRoutes(&userHandler)

	groupHandler := group.GroupHandler{
		Service: &group.GroupService{
			Repo: &group.GroupRepository{
				DB: dbConn,
			},
		},
	}
	api.SetupGroupRoutes(&groupHandler)
	log.Println("Server is listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
