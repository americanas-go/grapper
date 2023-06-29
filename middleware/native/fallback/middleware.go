package fallback

import (
	"github.com/americanas-go/grapper"
)

type anyErrorMiddleware[R any] struct {
}

func (c *anyErrorMiddleware[R]) Exec(ctx *grapper.AnyErrorContext[R], exec grapper.AnyErrorExecFunc[R], returnFunc grapper.AnyErrorReturnFunc[R]) (R, error) {
	r, err := ctx.Next(exec, returnFunc)
	return returnFunc(ctx.GetContext(), r, err)
}

func NewAnyErrorMiddleware[R any]() grapper.AnyErrorMiddleware[R] {
	return &anyErrorMiddleware[R]{}
}
