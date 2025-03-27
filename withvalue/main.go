// 値の受け渡し
package main

import (
	"context"
	"fmt"
)

func printUser(ctx context.Context) {
	userID := ctx.Value("userID")
	if userID != nil {
		fmt.Println("ユーザーID:", userID)
	} else {
		fmt.Println("ユーザーIDが見つかりません")
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)
	printUser(ctx) 
}
