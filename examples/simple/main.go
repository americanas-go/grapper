package main

import (
	"context"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/americanas-go/grapper"
	h "github.com/americanas-go/grapper/middleware/contrib/afex/hystrix-go.v0"
	"github.com/americanas-go/grapper/middleware/contrib/americanas-go/log.v1"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {

	ctx := context.Background()

	logrus.NewLogger()

	var r string
	var err error

	middlewares := []grapper.Middleware[string]{
		log.New[string](),
		h.NewWithConfig[string]("XPTO", hystrix.CommandConfig{
			Timeout:                10,
			MaxConcurrentRequests:  6000,
			RequestVolumeThreshold: 6000,
			SleepWindow:            10,
			ErrorPercentThreshold:  2,
		}),
	}

	wrp := grapper.New[string]("example", middlewares...)

	r, err = wrp.Exec(ctx,
		func(ctx context.Context) (string, error) {
			return "string", nil
		}, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(r)
}
