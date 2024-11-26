package auth

import (
	"gf-vue3-admin/api/auth"
	"gf-vue3-admin/internal/service/register"
)

func init() {
	register.Register("auth", NewV1())
}

type ControllerV1 struct{}

func NewV1() auth.IAuthV1 {
	return &ControllerV1{}
}
