package info

import "context"

type enemy int

const KeyCaojun = enemy(0)

func WithEnemyContext(ctx context.Context, number int) context.Context {
	return context.WithValue(ctx, KeyCaojun, number)
}

func FromEnemyContext(ctx context.Context) int {
	val := ctx.Value(KeyCaojun)
	n, ok := val.(int)
	if !ok {
		return 0
	}
	return n
}
