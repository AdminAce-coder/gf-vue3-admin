// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package apitest

import (
	"gf-vue3-admin/api/apitest"
	"gf-vue3-admin/internal/service/register"
)

func init() {
	register.Register("apitest", NewV1())
}

type ControllerV1 struct{}

func NewV1() apitest.IApitestV1 {
	return &ControllerV1{}
}
