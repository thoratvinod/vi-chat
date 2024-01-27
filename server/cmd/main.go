package main

import (
	"log"
	"net/http"

	"github.com/thoratvinod/vi-chat/inithandlers"
	"github.com/thoratvinod/vi-chat/usermgmt"
)

func main() {
	dbConn, err := inithandlers.InitDatabase()
	if err != nil {
		log.Fatalf("DB connection attempt failed, err=%+v", err)
	}

	inithandlers.SetupRoutes(dbConn)
	usermgmt.SetupRoutes(dbConn)
	log.Println("Server is listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
