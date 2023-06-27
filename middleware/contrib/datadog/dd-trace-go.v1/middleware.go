package datadog

import (
	"github.com/americanas-go/grapper"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type anyErrorMiddleware[R any] struct {
	name string
	tp   string
}

func (c *anyErrorMiddleware[R]) Exec(ctx *grapper.AnyErrorContext[R], exec grapper.AnyErrorExecFunc[R], returnFunc grapper.AnyErrorReturnFunc[R]) (r R, err error) {
	span, sctx := tracer.StartSpanFromContext(
		ctx.GetContext(),
		c.name,
		tracer.SpanType(c.tp),
	)
	defer span.Finish()

	ctx.SetContext(sctx)

	return ctx.Next(exec, returnFunc)
}

func NewAnyErrorMiddleware[R any](name string, tp string) grapper.AnyErrorMiddleware[R] {
	return &anyErrorMiddleware[R]{name: name, tp: tp}
}

type anyMiddleware[R any] struct {
	name string
	tp   string
}

func (c *anyMiddleware[R]) Exec(ctx *grapper.AnyContext[R], exec grapper.AnyExecFunc[R], returnFunc grapper.AnyReturnFunc[R]) (r R) {
	span, sctx := tracer.StartSpanFromContext(
		ctx.GetContext(),
		c.name,
		tracer.SpanType(c.tp),
	)
	defer span.Finish()

	ctx.SetContext(sctx)

	return ctx.Next(exec, returnFunc)
}

func NewAnyMiddleware[R any](name string, tp string) grapper.AnyErrorMiddleware[R] {
	return &anyErrorMiddleware[R]{name: name, tp: tp}
}

type errorMiddleware struct {
	name string
	tp   string
}

func (c *errorMiddleware) Exec(ctx *grapper.ErrorContext, exec grapper.ErrorExecFunc, returnFunc grapper.ErrorReturnFunc) (err error) {
	span, sctx := tracer.StartSpanFromContext(
		ctx.GetContext(),
		c.name,
		tracer.SpanType(c.tp),
	)
	defer span.Finish()

	ctx.SetContext(sctx)

	return ctx.Next(exec, returnFunc)
}

func NewErrorMiddleware(name string, tp string) grapper.ErrorMiddleware {
	return &errorMiddleware{name: name, tp: tp}
}
