package grapper

import "context"

type ExecFunc[R any] func(context.Context) (R, error)
type ReturnFunc[R any] func(context.Context, error) (R, error)

type Middleware[R any] interface {
	Exec(ctx *Context[R], exec ExecFunc[R], returnFunc ReturnFunc[R]) (R, error)
}
