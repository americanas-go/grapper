package log

import (
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/log"
)

type middleware[R any] struct{}

func (c *middleware[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], fallbackFunc grapper.FallbackFunc[R]) (R, error) {
	l := log.FromContext(ctx.GetContext())
	l.Tracef("executing wrapper %s", ctx.GetName())
	defer l.Tracef("wrapper %s executed", ctx.GetName())
	return ctx.Next(exec, fallbackFunc)
}

func New[R any]() grapper.Middleware[R] {
	return &middleware[R]{}
}
