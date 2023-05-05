package model

import "github.com/gogf/gf/v2/frame/g"

type OrderCreateReq struct {
	g.Meta  `json:"-" path:"/create-order" method:"post"`
	OrderId string `json:"orderId"`
}

type OrderCreateRes struct {
}

type OrderOneReq struct {
	g.Meta `json:"-" path:"/get-order" method:"get"`
}

type OrderOneRes struct {
	OrderId string `json:"orderId"`
}
