package hystrix

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/americanas-go/errors"
	"github.com/americanas-go/grapper"
	h "github.com/americanas-go/ignite/afex/hystrix-go.v0"
	"github.com/americanas-go/log"
)

type Hystrix[R any] struct {
	name string
}

func (c *Hystrix[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], returnFunc grapper.ReturnFunc[R]) (r R, err error) {

	if err = hystrix.DoC(ctx.GetContext(), c.name,
		func(ctxx context.Context) error {
			r, err = ctx.Next(ctx, exec, returnFunc)
			if err != nil {
				return err
			}
			return nil
		},
		func(ctxx context.Context, errr error) error {
			r, err = returnFunc(ctxx, errr)
			return err
		}); err != nil {
		return r, errors.Annotate(err, "error during execute hystrix circuit breaker")
	}

	return r, err
}

func New[R any](name string) grapper.Middleware[R] {
	if err := h.ConfigureCommand(name); err != nil {
		log.Error(err.Error())
	}
	return &Hystrix[R]{name: name}
}
