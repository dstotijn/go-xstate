package xstate

import "context"

type InvokeConfig struct {
	ID          string
	Src         Invoker
	OnDone      []Transition
	OnError     []Transition
	AutoForward bool
	Data        map[string]func(context Data, event Event) interface{}
}

type Invoker interface {
	Invoke(ctx context.Context, event Event)
}

type InvokerFunc func(ctx context.Context, event Event) interface{}

func (f InvokerFunc) Invoke(ctx context.Context, event Event) interface{} {
	return f(ctx, event)
}
