package group

import (
	"encoding/json"
	"fmt"
	"net/http"

	httpCommon "github.com/thoratvinod/vi-chat/server/pkg/common/http"
)

type GroupHandler struct {
	Service *GroupService
}

func (handler *GroupHandler) SaveGroup(w http.ResponseWriter, r *http.Request) {
	grp := Group{}
	err := json.NewDecoder(r.Body).Decode(&grp)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
		return
	}
	err = handler.Service.SaveGroup(&grp)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("save group request failed ; %v", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (handler *GroupHandler) AddMembers(w http.ResponseWriter, r *http.Request) {
	req := struct {
		MemberIDs []uint `json:"memberIDs"`
		GroupID   uint   `json:"groupID"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
		return
	}
	err = handler.Service.AddMembers(req.GroupID, req.MemberIDs)
	if err != nil {
		httpCommon.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}

func (handler *GroupHandler) RemoveMembers(w http.ResponseWriter, r *http.Request) {
	req := struct {
		MemberIDs []uint `json:"memberIDs"`
		GroupID   uint   `json:"groupID"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		httpCommon.HandleError(w, fmt.Errorf("invalid request body ; %v", err.Error()), http.StatusBadRequest)
		return
	}
	err = handler.Service.RemoveMembers(req.GroupID, req.MemberIDs)
	if err != nil {
		httpCommon.HandleError(w, err, http.StatusInternalServerError)
		return
	}
}
