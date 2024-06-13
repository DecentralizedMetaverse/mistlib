package social

import (
	"encoding/json"
	"net/http"
)

type GetWorldListRequest struct {
	UserID string `json:"user_id"`
}

type GetWorldListResponse struct {
	Worlds []string `json:"worlds"`
}

func GetWorldList(w http.ResponseWriter, r *http.Request) {
	var req GetWorldListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	worlds := []string{"world1", "world2", "world3"}
	res := GetWorldListResponse{Worlds: worlds}
	json.NewEncoder(w).Encode(res)
}
