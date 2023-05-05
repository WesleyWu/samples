package main

import (
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/samples/hello-go-server/middleware"
	"github.com/dapr/samples/hello-go-server/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.ResponseJsonWrapper)
		group.Bind(
			service.NewOrderService(client, "statestore"),
		)
	})
	s.Run()
}
