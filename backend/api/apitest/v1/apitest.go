package v1

import "github.com/gogf/gf/v2/frame/g"

type ApiDemoTaReq struct {
	g.Meta `path:"/apidemota" method:"put" tags:"apitest" dc:""`
}

type ApiDemoTaRes struct {
}
