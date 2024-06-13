package world_creation

import (
	"fmt"
)

func PutObject(objectID, userID, worldID string) error {
	// オブジェクト配置処理
	fmt.Println("Object placed:", objectID, "by user:", userID, "in world:", worldID)
	return nil
}

func AddComponent(objectID, componentID string) error {
	// コンポーネント追加処理
	fmt.Println("Component added:", componentID, "to object:", objectID)
	return nil
}
