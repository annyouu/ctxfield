
// ダメな例
// package main

// import (
// 	"context"
// )

// // NG: 構造体にcontextを直接保持する
// type S struct {
// 	ctx context.Context
// }

// // NG: コンテキストを使う関数
// func (s *S) DoSomething() {
// 	select {
// 	case <-s.ctx.Done():
// 		println("コンテキストがキャンセルされました")
// 	default:
// 		println("処理を継続...")
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	s := S{ctx: ctx} // NG: 構造体にコンテキストを保持
// 	go s.DoSomething()

// 	cancel()
// }


// ---------------------------------------
// 推奨コード(コンテキストを関数の引数にする)
package main

import (
	"context"
	"fmt"
	"time"
)

func DoSomething(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("コンテキストがキャンセルされた")
	default:
		fmt.Println("処理を継続...")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go DoSomething(ctx)

	time.Sleep(1 * time.Second)
	cancel()
}
