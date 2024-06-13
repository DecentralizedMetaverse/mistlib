package world_creation

import (
	"encoding/json"
	"net/http"
)

type PutObjectRequest struct {
	ObjectID string `json:"object_id"`
	UserID   string `json:"user_id"`
	WorldID  string `json:"world_id"`
}

func PutObject(w http.ResponseWriter, r *http.Request) {
	var req PutObjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
