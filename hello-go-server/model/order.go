package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type OrderCreateReq struct {
	g.Meta    `json:"-" path:"/create-order" method:"post"`
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type OrderCreateRes struct {
}

type OrderOneReq struct {
	g.Meta `json:"-" path:"/get-order" method:"get"`
	Id     string `json:"id"`
}

type OrderOneRes struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type OrderListReq struct {
	g.Meta `json:"-" path:"/list-order" method:"post"`
	Query  string `json:"query"`
}

type OrderItem struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

type OrderListRes struct {
	Items []*OrderItem `json:"items"`
}
