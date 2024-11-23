package v1

import "github.com/gogf/gf/v2/frame/g"

// 系统管理相关API
type Parameter struct {
	ParameterName string `json:"parametername"` // 参数名
	DataType      string `json:"datatype"`      // 数据类型
	Required      bool   `json:"required"`
	Description   string `json:"description"` // 描述
}

type CreateApiReq struct {
	g.Meta           `path:"/createapi" method:"post" tags:"login" dc:"新增API"`
	ApiPath          string      `json:"apiPath"`
	Isauthentication bool        `json:"isauthentication"` // 是否鉴权
	ApiVersion       string      `json:"apiVersion"`
	Method           string      `json:"method"`
	ApiGroup         string      `json:"apiGroup"`
	Description      string      `json:"description"`
	Parameters       []Parameter `json:"parameters"` // 参数
}
type CreateApiRes struct {
}
