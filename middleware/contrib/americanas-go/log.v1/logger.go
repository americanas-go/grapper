package log

import (
	"github.com/americanas-go/grapper"
	"github.com/americanas-go/log"
)

type logger[R any] struct{}

func (c *logger[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], returnFunc grapper.ReturnFunc[R]) (R, error) {
	l := log.FromContext(ctx.GetContext())
	l.Tracef("executing wrapper %s", ctx.GetName())
	defer l.Debugf("wrapper %s executed", ctx.GetName())
	return ctx.Next(ctx, exec, returnFunc)
}

func New[R any]() grapper.Middleware[R] {
	return &logger[R]{}
}
