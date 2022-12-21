package main

import (
	"context"
	"fmt"

	"github.com/americanas-go/config"
	"github.com/americanas-go/grapper"
	h "github.com/americanas-go/grapper/middleware/contrib/americanas-go/ignite.v1/afex/hystrix-go.v0"
	"github.com/americanas-go/grapper/middleware/contrib/americanas-go/log.v1"
	"github.com/americanas-go/ignite/afex/hystrix-go.v0"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func init() {
	hystrix.CommandConfigAdd("XPTO")
}

func main() {

	ctx := context.Background()

	logrus.NewLogger()
	config.Load()

	var r string
	var err error

	wrp := grapper.New[string](log.New[string](), h.New[string]("XPTO"))

	r, err = wrp.Exec(ctx,
		func(ctx context.Context) (string, error) {
			return "string", nil
		}, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(r)
}
