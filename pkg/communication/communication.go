package communication

import (
	"fmt"
)

func SendMessage(fromUserID, toUserID, message string) error {
	// メッセージ送信処理
	fmt.Println("Message sent from:", fromUserID, "to:", toUserID, "message:", message)
	return nil
}

func CreateGroup(groupName, userID string) error {
	// グループ作成処理
	fmt.Println("Group created:", groupName, "by user:", userID)
	return nil
}
