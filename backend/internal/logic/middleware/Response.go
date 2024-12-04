package middleware

import (
	"gf-vue3-admin/internal/service"
	"gf-vue3-admin/utility/docode"
	"net/http"

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
		msg        string
		rcode      int
		httpStatus = http.StatusOK // 默认状态码为 200
		res        = r.GetHandlerResponse()
		err        = r.GetError()
	)
	if err != nil {
		if e, ok := err.(*docode.ErrorWithCode); ok {
			rcode = e.Code()
			msg = e.Error()

			isCodeSet := rcode != 0 // 检查是否传入了有效的状态码
			if isCodeSet {
				httpStatus = rcode
			} else {
				httpStatus = http.StatusBadRequest
			}
			// 返回错误的
			// r.Response.WriteStatus(httpStatus)
			r.Response.WriteJson(Response{
				Code:    rcode,
				Message: msg,
				Data:    res,
			})
			r.Response.WriteHeader(httpStatus) // 使用 WriteHeader 替代 WriteStatus
		}
	} else {
		//rcode = code2.SUCCESS
		msg = "success"
		// r.Response.WriteStatus(httpStatus) // 成功状态使用 200
		r.Response.WriteJson(Response{
			Code:    0,
			Message: msg,
			Data:    res,
		})
		r.Response.WriteHeader(httpStatus) // 使用 WriteHeader 替代 WriteStatus
	}
}
