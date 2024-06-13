package communication

import (
	"encoding/json"
	"net/http"
)

type CreateGroupRequest struct {
	GroupName string `json:"group_name"`
	UserID    string `json:"user_id"`
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var req CreateGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
