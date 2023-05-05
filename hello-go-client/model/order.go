package model

import "github.com/gogf/gf/v2/frame/g"

type OrderCreateReq struct {
	g.Meta  `json:"-" path:"/createorder" method:"post"`
	OrderId string `json:"orderId"`
}

type OrderCreateRes struct {
}

type OrderOneReq struct {
	g.Meta `json:"-" path:"/getorder" method:"get"`
}

type OrderOneRes struct {
	OrderId string `json:"orderId"`
}
