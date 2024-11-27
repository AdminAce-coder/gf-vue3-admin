// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package apitest

import (
	"context"

	"gf-vue3-admin/api/apitest/v1"
)

type IApitestV1 interface {
	ApiDemoTa(ctx context.Context, req *v1.ApiDemoTaReq) (res *v1.ApiDemoTaRes, err error)
}
