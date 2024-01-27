package common

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var contentType = "application/json"

type Controller interface {
	SetupRoutes()
}

func RegisterRoute(method, pattern string, endpoint func(http.ResponseWriter, *http.Request) (any, ErrorResponse)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		if r.Method != method {
			return
		}
		resp, err := endpoint(w, r)
		if err != nil {
			WriteErrorJSON(w, err)
			return
		}
		WriteJSON(w, resp)
		log.Printf("method=%v path=%v timeElapsed=%vms", method, pattern, time.Now().Sub(startTime).Milliseconds())
	})
}

func WriteJSON(w http.ResponseWriter, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		WriteErrorJSON(w, NewErrorResponse(DefaultErrorCode, DefaultErrorMessage))
		return
	}
	w.Header().Set("Content-Type", contentType)
	w.Write(b)
}

func WriteErrorJSON(w http.ResponseWriter, err ErrorResponse) {
	w.Header().Set("Content-Type", contentType)
	status := GetStatusCode(err.GetErrorCode())
	w.WriteHeader(status)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}
