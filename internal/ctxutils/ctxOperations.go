package ctxutil

import (
	"context"
	"fmt"
)

func CreateContext(ctx context.Context, key ContextKey, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func AddContext(ctx context.Context, key1 ContextKey, value1 any, key2 ContextKey, value2 any) context.Context {
	c := CreateContext(ctx, key1, value1)
	return CreateContext(c, key2, value2)
}

func GetContext(ctx context.Context, key ContextKey) (any, error) {
	val := ctx.Value(key)
	if val == nil {
		return nil, fmt.Errorf("Context key not found")
	}

	return ctx.Value(key), nil
}
