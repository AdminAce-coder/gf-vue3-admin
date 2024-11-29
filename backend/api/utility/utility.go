// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package utility

import (
	"context"

	"gf-vue3-admin/api/utility/v1"
)

type IUtilityV1 interface {
	Utility_Create_Oss(ctx context.Context, req *v1.Utility_Create_OssReq) (res *v1.Utility_Create_OssRes, err error)
}
