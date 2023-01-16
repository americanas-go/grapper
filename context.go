package grapper

import "context"

type Context[R any] struct {
	ctx   context.Context
	m     []Middleware[R]
	name  string
	index int
	id    string
}

func (c *Context[R]) GetName() string {
	return c.name
}

func (c *Context[R]) GetContext() context.Context {
	return c.ctx
}

func (c *Context[R]) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *Context[R]) SetID(id string) {
	c.id = id
}

func (c *Context[R]) GetID() string {
	return c.id
}

func (c *Context[R]) Next(exec ExecFunc[R], fallback FallbackFunc[R]) (R, error) {
	if m := c.getNext(); m != nil {
		return m.Exec(c, exec, fallback)
	}
	return exec(c.GetContext())
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

func NewContext[R any](name string, id string, m ...Middleware[R]) *Context[R] {
	return &Context[R]{m: m, name: name, id: id}
}
