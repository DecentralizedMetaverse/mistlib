package social

import (
	"encoding/json"
	"net/http"
)

type GetUserListRequest struct {
	WorldID string `json:"world_id"`
}

type GetUserListResponse struct {
	Users []string `json:"users"`
}

func GetUserList(w http.ResponseWriter, r *http.Request) {
	var req GetUserListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users := []string{"user1", "user2", "user3"}
	res := GetUserListResponse{Users: users}
	json.NewEncoder(w).Encode(res)
}
