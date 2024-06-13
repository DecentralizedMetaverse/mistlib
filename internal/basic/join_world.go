package basic

import (
	"encoding/json"
	"net/http"
)

type JoinWorldRequest struct {
	UserID  string `json:"user_id"`
	WorldID string `json:"world_id"`
}

func JoinWorld(w http.ResponseWriter, r *http.Request) {
	var req JoinWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
