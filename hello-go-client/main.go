package main

import (
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/samples/hello-go-client/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	req := &model.OrderOneReq{}
	out, err := client.InvokeMethodWithCustomContent(ctx, "hello-go-server", "get-order", "GET", "application/json", req)
	if err != nil {
		panic(err)
	}
	g.Log().Infof(ctx, "out: %v", string(out))
}
