package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("処理完了")
	case <-ctx.Done():
		fmt.Println("処理がタイムアウトしました:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel() // キャンセル後にリソースを解放

	go doWork(ctx)

	time.Sleep(3 * time.Second) // 3秒待機
	fmt.Println("メイン処理終了")
}