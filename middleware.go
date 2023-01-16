package grapper

import "context"

type ExecFunc[R any] func(context.Context) (R, error)
type FallbackFunc[R any] func(context.Context, R, error) (R, error)

type Middleware[R any] interface {
	Exec(ctx *Context[R], exec ExecFunc[R], returnFunc FallbackFunc[R]) (R, error)
}
