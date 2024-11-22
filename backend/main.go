package main

import (
	"gf-vue3-admin/internal/cmd"
	_ "gf-vue3-admin/internal/logic"
	_ "gf-vue3-admin/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
