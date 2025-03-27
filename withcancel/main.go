package main

import (
	"context"
	"fmt"
	"time"
)

// 　長時間かかる処理を模擬
func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("処理がキャンセルされました")
			return
		default:
			fmt.Println("処理中..")
			time.Sleep(1 * time.Second)
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go doWork(ctx) // キャンセル可能なコンテキストを渡す

	time.Sleep(3 * time.Second)
	fmt.Println("キャンセルを実行")
	cancel() // キャンセル信号を送る

	time.Sleep(1 * time.Second) // 少し待ってから終了
}
