package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/scripts/perftest"
)

const (
	ConcurrentConnections = 1000
	ConcurrentMessages    = 10000
)

const (
	baseURL = "http://localhost:8080"
)

var JWTTokenMap = map[uint]string{}
var ConnMap = map[uint]*websocket.Conn{}

func signup() {
	for i := 1; i <= ConcurrentConnections; i++ {
		body := []byte(fmt.Sprintf(`{
			{
				"firstNmae": "first-name-%v",
				"lastName": "last-name-%v",
				"email": "email%v@gmail.com",
				"password": "test-%v"
			}
		}`, i, i, i, i))

		r, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("%v/user/signup", baseURL),
			bytes.NewBuffer(body),
		)
		if err != nil {
			log.Printf("Got error while creating request for signup, userID=%v, err=%+v", i, err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		_, err = client.Do(r)
		if err != nil {
			log.Printf("Got error while firing request for signup, userID=%v, err=%+v", i, err.Error())
			continue
		}
	}
}

func login() {
	for i := uint(1); i <= ConcurrentConnections; i++ {

		body := []byte(fmt.Sprintf(`{
			{
				"email": "email%v@gmail.com",
				"password": "test-%v"
			}
		}`, i, i))

		r, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("%v/user/login", baseURL),
			bytes.NewBuffer(body),
		)
		if err != nil {
			log.Printf("Got error while creating request for login, userID=%v, err=%+v", i, err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json")
		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			log.Printf("Got error while firing request for login, userID=%v, err=%+v", i, err.Error())
			continue
		}

		token := struct {
			JWTToken string `json:"jwtToken"`
		}{}

		err = json.NewDecoder(res.Body).Decode(&token)
		if err != nil {
			log.Printf("Got error while decoding response for login, userID=%v, err=%+v", i, err.Error())
			continue
		}
		JWTTokenMap[i] = token.JWTToken
		res.Body.Close()
	}
}

func main() {

	signup()
	login()

	connHandler := perftest.NewClientConnectionHandler(JWTTokenMap)
	connHandler.EstablishConnections(baseURL)

}
