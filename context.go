package grapper

import "context"

type Context[R any] struct {
	ctx   context.Context
	m     []Middleware[R]
	index int
}

func (c *Context[R]) GetContext() context.Context {
	return c.ctx
}

func (c *Context[R]) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *Context[R]) Next(ctx *Context[R], exec ExecFunc[R], returnFunc ReturnFunc[R]) (R, error) {
	if m := c.getNext(); m != nil {
		return m.Exec(ctx, exec, returnFunc)
	}
	return exec(ctx.GetContext())
}

func (c *Context[R]) hasNext() bool {
	if c.index < len(c.m) {
		return true
	}
	return false

}
func (c *Context[R]) getNext() Middleware[R] {
	if c.hasNext() {
		m := c.m[c.index]
		c.index++
		return m
	}
	return nil
}

func NewContext[R any](m ...Middleware[R]) *Context[R] {
	return &Context[R]{m: m}
}
