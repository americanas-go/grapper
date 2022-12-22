package datadog

import (
	"github.com/americanas-go/grapper"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type span[R any] struct {
	name string
	tp   string
}

func (c *span[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], returnFunc grapper.ReturnFunc[R]) (R, error) {
	span, sctx := tracer.StartSpanFromContext(
		ctx.GetContext(),
		c.name,
		tracer.SpanType(c.tp),
	)
	defer span.Finish()

	ctx.SetContext(sctx)

	return ctx.Next(ctx, exec, returnFunc)
}

func NewSpan[R any](name string, tp string) grapper.Middleware[R] {
	return &span[R]{name: name, tp: tp}
}
