package log

import (
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/log"
)

type Logger[R any] struct {
}

func (c *Logger[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], returnFunc grapper.ReturnFunc[R]) (R, error) {
	log.FromContext(ctx.GetContext()).Infof("executing method")
	return ctx.Next(ctx, exec, returnFunc)
}

func New[R any]() grapper.Middleware[R] {
	return &Logger[R]{}
}
