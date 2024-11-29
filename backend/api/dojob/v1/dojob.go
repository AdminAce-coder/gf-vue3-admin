package v1

import "github.com/gogf/gf/v2/frame/g"

type CommandInputReq struct {
	g.Meta `path:"/commandinput" method:"post" tags:"dojob" dc:""`
	// 命令
	command string `v:"required" dc:"命令"`
}

type CommandInputRes struct {
}
