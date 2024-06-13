package social

import (
	"encoding/json"
	"net/http"
)

type GetFriendListRequest struct {
	UserID string `json:"user_id"`
}

type GetFriendListResponse struct {
	Friends []string `json:"friends"`
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	var req GetFriendListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	friends := []string{"friend1", "friend2", "friend3"}
	res := GetFriendListResponse{Friends: friends}
	json.NewEncoder(w).Encode(res)
}
