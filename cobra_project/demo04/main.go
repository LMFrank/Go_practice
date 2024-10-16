package main

import (
	"context"
	"demo04/info"
	"fmt"
)

// Output:
// 刘备: 曹操来了 9 万人
// 关羽(1) <-: 曹操来了 9 万人
// 关羽(2) ->: 曹操来了 90 万人
// 张飞: 曹操来了 90 万人
func main() {
	ctx := context.Background()

	Liubei(ctx, 9)
}

func Liubei(ctx context.Context, number int) {
	ctx = info.WithEnemyContext(ctx, number)
	fmt.Printf("刘备： 曹操来了 %d 万人\n", number)
	Guanyu(ctx)
}

func Guanyu(ctx context.Context) {
	n := info.FromEnemyContext(ctx)
	fmt.Printf("关羽（1）<-: 曹操来了 %d 万人\n", n)

	if n%2 == 1 {
		n = n * 10
		ctx = info.WithEnemyContext(ctx, n)
	}

	fmt.Printf("关羽（2）->: 曹操来了 %d 万人\n", n)
	Zhangfei(ctx)
}

func Zhangfei(ctx context.Context) {
	n := info.FromEnemyContext(ctx)
	fmt.Printf("张飞: 曹操来了 %d 万人\n", n)
}
