package log

import (
	"context"

	"github.com/americanas-go/cache"
	"github.com/americanas-go/grapper"
)

type middleware[R any] struct {
	manager *cache.Manager[R]
	opts    []cache.OptionSet
}

func (m *middleware[R]) Exec(c *grapper.Context[R], exec grapper.ExecFunc[R], fallbackFunc grapper.FallbackFunc[R]) (R, error) {
	return m.manager.GetOrSet(c.GetContext(), c.GetID(), func(ctx context.Context) (R, error) {
		return c.Next(exec, fallbackFunc)
	}, m.opts...)
}

func New[R any](manager *cache.Manager[R], opts ...cache.OptionSet) grapper.Middleware[R] {
	return &middleware[R]{manager: manager, opts: opts}
}
