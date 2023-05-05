package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ResponseJsonWrapper(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器开小差了，请稍后再试吧！")
	} else {
		if err := r.GetError(); err != nil {
			r.Response.ClearBuffer()
			r.Response.Status = 500
			r.Response.WriteJsonExit(err)
		} else {
			handlerRes := r.GetHandlerResponse()
			r.Response.WriteJsonExit(handlerRes)
		}
	}
}
