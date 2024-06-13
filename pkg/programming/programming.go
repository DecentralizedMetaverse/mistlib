package programming

import (
	"fmt"
)

func CallRPC(method, params string) (string, error) {
	// RPCコール処理
	result := fmt.Sprintf("Method: %s, Params: %s", method, params)
	fmt.Println("RPC called:", result)
	return result, nil
}
