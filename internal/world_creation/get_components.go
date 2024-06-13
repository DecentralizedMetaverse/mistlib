package world_creation

import (
	"encoding/json"
	"net/http"
)

type GetComponentsRequest struct {
	ObjectID string `json:"object_id"`
}

type GetComponentsResponse struct {
	Components []string `json:"components"`
}

func GetComponents(w http.ResponseWriter, r *http.Request) {
	var req GetComponentsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	components := []string{"component1", "component2", "component3"}
	res := GetComponentsResponse{Components: components}
	json.NewEncoder(w).Encode(res)
}
