package social

import (
	"fmt"
)

func AddFriend(userID, friendID string) error {
	// フレンド追加処理
	fmt.Println("Friend added:", friendID, "for user:", userID)
	return nil
}

func RemoveFriend(userID, friendID string) error {
	// フレンド削除処理
	fmt.Println("Friend removed:", friendID, "for user:", userID)
	return nil
}
