package world_creation

import (
	"encoding/json"
	"net/http"
)

type RemoveComponentRequest struct {
	ObjectID    string `json:"object_id"`
	ComponentID string `json:"component_id"`
}

func RemoveComponent(w http.ResponseWriter, r *http.Request) {
	var req RemoveComponentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
