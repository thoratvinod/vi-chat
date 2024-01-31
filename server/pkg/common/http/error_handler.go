package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func HandleError(w http.ResponseWriter, err error, statusCode int) {
	log.Println(err)
	w.WriteHeader(statusCode)
	errResp := ErrorResponse{Message: err.Error()}
	json.NewEncoder(w).Encode(errResp)
}
