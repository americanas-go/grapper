package datadog

import (
	"github.com/americanas-go/grapper"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type middleware[R any] struct {
	name string
	tp   string
}

func (c *middleware[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], fallbackFunc grapper.FallbackFunc[R]) (R, error) {
	span, sctx := tracer.StartSpanFromContext(
		ctx.GetContext(),
		c.name,
		tracer.SpanType(c.tp),
	)
	defer span.Finish()

	ctx.SetContext(sctx)

	return ctx.Next(exec, fallbackFunc)
}

func New[R any](name string, tp string) grapper.Middleware[R] {
	return &middleware[R]{name: name, tp: tp}
}
