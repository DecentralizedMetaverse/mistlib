package communication

import (
	"encoding/json"
	"net/http"
)

type LeaveGroupRequest struct {
	GroupID string `json:"group_id"`
	UserID  string `json:"user_id"`
}

func LeaveGroup(w http.ResponseWriter, r *http.Request) {
	var req LeaveGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
