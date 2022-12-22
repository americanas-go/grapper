package grapper

import (
	"context"
)

type Wrapper[R any] struct {
	m    []Middleware[R]
	name string
}

func New[R any](name string, m ...Middleware[R]) *Wrapper[R] {
	return &Wrapper[R]{m: m, name: name}
}

func (w *Wrapper[R]) Exec(ctx context.Context, exec ExecFunc[R], returnFunc ReturnFunc[R]) (r R, err error) {
	c := NewContext(w.name, w.m...)
	c.SetContext(ctx)
	return c.Next(c, exec, returnFunc)
}
