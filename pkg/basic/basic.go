package basic

import (
	"fmt"
)

func ConnectUser(userID, password string) error {
	// ユーザーの接続処理
	fmt.Println("User connected:", userID)
	return nil
}

func DisconnectUser(userID string) error {
	// ユーザーの切断処理
	fmt.Println("User disconnected:", userID)
	return nil
}
