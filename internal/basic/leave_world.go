package basic

import (
	"encoding/json"
	"net/http"
)

type LeaveWorldRequest struct {
	UserID  string `json:"user_id"`
	WorldID string `json:"world_id"`
}

func LeaveWorld(w http.ResponseWriter, r *http.Request) {
	var req LeaveWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
