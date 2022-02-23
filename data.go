package xstate

import "context"

type Data map[string]interface{}

type ctxKey int

var dataCtxKey ctxKey

func NewContextWithData(ctx context.Context, data Data) context.Context {
	return context.WithValue(ctx, dataCtxKey, data)
}

func DataFromContext(ctx context.Context) (Data, bool) {
	data, ok := ctx.Value(dataCtxKey).(Data)
	return data, ok
}
