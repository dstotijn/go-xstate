package xstate

import "context"

type Action struct {
	Type string
	Exec ActionFn
}

type ActionFn func(context.Context, Event, Meta)

func Assign(assignMap map[string]func(context.Context, Event, Meta) interface{}) ActionFn {
	return func(ctx context.Context, event Event, meta Meta) {
		data, ok := DataFromContext(ctx)
		if !ok {
			return
		}

		for key, fn := range assignMap {
			x := fn(ctx, event, meta)
			data[key] = x
		}
	}
}
