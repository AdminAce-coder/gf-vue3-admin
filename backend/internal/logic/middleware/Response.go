package middleware

import (
	code2 "gf-vue3-admin/internal/consts/code/currency"
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/docode"

	"github.com/gogf/gf/v2/net/ghttp"
)

// @ 返回统一结构的中间件
type Response struct {
	Code    int         `json:"code"    dc:"状态码"`
	Message string      `json:"message" dc:"消息提示"`
	Data    interface{} `json:"data"    dc:"执行结果"`
}

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(new())
}

func new() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Returndata(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg   string
		rcode int
		res   = r.GetHandlerResponse()
		err   = r.GetError()
	)
	if err != nil {
		if e, ok := err.(*docode.ErrorWithCode); ok {
			rcode = e.Code()
			msg = e.Error()
			// 返回错误的
			r.Response.WriteJson(Response{
				Code:    rcode,
				Message: msg,
				Data:    res,
			})
		}
	} else {
		rcode = code2.SUCCESS
		msg = "success"
		r.Response.WriteJson(Response{
			Code:    rcode,
			Message: msg,
			Data:    res,
		})
	}

}
