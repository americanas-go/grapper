package main

import (
	"context"
	"fmt"

	"github.com/americanas-go/grapper"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

type CustomMiddleware[R any] struct{}

func (c *CustomMiddleware[R]) Exec(ctx *grapper.Context[R], exec grapper.ExecFunc[R], fallbackFunc grapper.FallbackFunc[R]) (R, error) {
	fmt.Println("my custom middleware")
	return ctx.Next(exec, fallbackFunc)
}

func NewCustomMiddleware[R any]() grapper.Middleware[R] {
	return &CustomMiddleware[R]{}
}

func main() {

	ctx := context.Background()

	logrus.NewLogger()

	var res string
	var err error

	middlewares := []grapper.Middleware[string]{
		NewCustomMiddleware[string](),
	}

	wrp := grapper.New[string]("example", middlewares...)

	res, err = wrp.Exec(ctx, "1",
		func(ctx context.Context) (string, error) {
			return "string", nil
		}, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
