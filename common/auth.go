package common

import (
	"fmt"
	"net/http"
	"strings"
)

func ExtractJWTToken(r *http.Request) (string, error) {
	// get auth jwt token
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return "", fmt.Errorf("authorization JWT token not provided")
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 || len(splitToken[1]) == 0 {
		return "", fmt.Errorf("Invalid JWT token received")
	}
	return splitToken[1], nil
}
