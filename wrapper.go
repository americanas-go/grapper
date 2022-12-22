package grapper

import (
	"context"
)

type Wrapper[R any] struct {
	m []Middleware[R]
}

func New[R any](m ...Middleware[R]) *Wrapper[R] {
	return &Wrapper[R]{m: m}
}

func (w *Wrapper[R]) Exec(ctx context.Context, exec ExecFunc[R], returnFunc ReturnFunc[R]) (r R, err error) {
	c := NewContext(w.m...)
	c.SetContext(ctx)
	return c.Next(c, exec, returnFunc)
}
