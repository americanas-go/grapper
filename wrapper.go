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

func (w *Wrapper[R]) Exec(ctx context.Context, id string, exec ExecFunc[R], returnFunc FallbackFunc[R]) (r R, err error) {
	c := NewContext(w.name, id, w.m...)
	c.SetContext(ctx)
	return c.Next(exec, returnFunc)
}
