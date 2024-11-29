// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dojob

import (
	"context"

	"gf-vue3-admin/api/dojob/v1"
)

type IDojobV1 interface {
	CommandInput(ctx context.Context, req *v1.CommandInputReq) (res *v1.CommandInputRes, err error)
}
