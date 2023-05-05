package service

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	p "github.com/dapr/samples/hello-go-server/model"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

func NewOrderService(client dapr.Client, stateStore string) IOrder {
	return &OrderImpl{
		client:     client,
		stateStore: stateStore,
	}
}

type IOrder interface {
	Create(ctx context.Context, req *p.OrderCreateReq) (*p.OrderCreateRes, error)
	Get(ctx context.Context, req *p.OrderOneReq) (*p.OrderOneRes, error)
}

type OrderImpl struct {
	client     dapr.Client
	stateStore string
}

func (s *OrderImpl) Create(ctx context.Context, req *p.OrderCreateReq) (*p.OrderCreateRes, error) {
	g.Log().Infof(ctx, "create req: %v", gjson.MustEncodeString(req))
	err := s.client.SaveState(ctx, s.stateStore, "order", gjson.MustEncode(req), nil)
	return &p.OrderCreateRes{}, err
}

func (s *OrderImpl) Get(ctx context.Context, req *p.OrderOneReq) (*p.OrderOneRes, error) {
	g.Log().Infof(ctx, "one req: %v", gjson.MustEncodeString(req))
	item, err := s.client.GetState(ctx, s.stateStore, "order", nil)
	if err != nil {
		return nil, err
	}
	value := item.Value
	g.Log().Infof(ctx, "state loaded: %v", string(value))
	result := &p.OrderOneRes{}
	if value != nil {
		err = gjson.DecodeTo(value, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
