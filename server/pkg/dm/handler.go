package dm

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	httpCommon "github.com/thoratvinod/vi-chat/server/pkg/common/http"
// )

// type DMHandler struct {
// 	Service *DMService
// }

// func (handler *DMHandler) SaveAndNotify(w http.ResponseWriter, r *http.Request) {
// 	var dm DirectMessage
// 	err := json.NewDecoder(r.Body).Decode(&dm)
// 	if err != nil {
// 		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.Service.SaveMessage(&dm)
// 	if err != nil {
// 		httpCommon.HandleError(w, fmt.Errorf("internal server error ; %v", err.Error()), http.StatusInternalServerError)
// 		return
// 	}
// }
