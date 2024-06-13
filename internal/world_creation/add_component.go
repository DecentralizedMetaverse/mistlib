package world_creation

import (
	"encoding/json"
	"net/http"
)

type AddComponentRequest struct {
	ObjectID    string `json:"object_id"`
	ComponentID string `json:"component_id"`
}

func AddComponent(w http.ResponseWriter, r *http.Request) {
	var req AddComponentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
