package social

import (
	"encoding/json"
	"net/http"
)

type RemoveFriendRequest struct {
	UserID   string `json:"user_id"`
	FriendID string `json:"friend_id"`
}

func RemoveFriend(w http.ResponseWriter, r *http.Request) {
	var req RemoveFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
