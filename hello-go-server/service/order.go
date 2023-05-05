package service

import (
	"context"

	dapr "github.com/dapr/go-sdk/client"
	p "github.com/dapr/samples/hello-go-server/model"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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
	List(ctx context.Context, req *p.OrderListReq) (*p.OrderListRes, error)
}

type OrderImpl struct {
	client     dapr.Client
	stateStore string
}

func (s *OrderImpl) Create(ctx context.Context, req *p.OrderCreateReq) (*p.OrderCreateRes, error) {
	g.Log().Infof(ctx, "create req: %v", gjson.MustEncodeString(req))
	id := "order|" + req.Id
	req.CreatedAt = gtime.Now()
	req.UpdatedAt = gtime.Now()
	err := s.client.SaveState(ctx, s.stateStore, id, gjson.MustEncode(req), nil)
	return &p.OrderCreateRes{}, err
}

func (s *OrderImpl) Get(ctx context.Context, req *p.OrderOneReq) (*p.OrderOneRes, error) {
	id := "order|" + req.Id
	item, err := s.client.GetState(ctx, s.stateStore, id, nil)
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

func (s *OrderImpl) List(ctx context.Context, req *p.OrderListReq) (*p.OrderListRes, error) {
	query := req.Query
	queryRes, err := s.client.QueryStateAlpha1(ctx, s.stateStore, query, nil)
	if err != nil {
		return nil, err
	}
	itemCount := len(queryRes.Results)
	items := make([]*p.OrderItem, itemCount)
	for i, item := range queryRes.Results {
		err = gjson.DecodeTo(item.Value, &items[i])
		if err != nil {
			return nil, err
		}
	}
	return &p.OrderListRes{Items: items}, nil
}
