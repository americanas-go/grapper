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

type Result struct {
	Code string
}

type FooService struct {
	wrapper *grapper.Wrapper[Result]
}

func NewFooService(wrapper *grapper.Wrapper[Result]) *FooService {
	return &FooService{wrapper: wrapper}
}

func (s *FooService) FooMethod(ctx context.Context) (Result, error) {
	return s.wrapper.Exec(ctx, "1", func(ctx context.Context) (Result, error) {
		return Result{Code: "XPTO"}, nil
	}, nil)
}

func main() {

	ctx := context.Background()

	logrus.NewLogger()

	var r Result
	var err error

	middlewares := []grapper.Middleware[Result]{
		log.New[Result](),
		h.NewWithConfig[Result]("XPTO", hystrix.CommandConfig{
			Timeout:                10,
			MaxConcurrentRequests:  6000,
			RequestVolumeThreshold: 6000,
			SleepWindow:            10,
			ErrorPercentThreshold:  2,
		}),
	}

	wrapper := grapper.New[Result]("example", middlewares...)

	foo := NewFooService(wrapper)
	r, err = foo.FooMethod(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Code)
}
