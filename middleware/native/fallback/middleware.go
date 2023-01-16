package fallback

import (
	"github.com/americanas-go/grapper"
)

type middleware[R any] struct {
}

func (c *middleware[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], fallback grapper.FallbackFunc[R]) (R, error) {
	r, err := ctx.Next(exec, fallback)
	return fallback(ctx.GetContext(), r, err)
}

func New[R any]() grapper.Middleware[R] {
	return &middleware[R]{}
}
