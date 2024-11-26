package register

import (
	"context"
	"gf-vue3-admin/utility"
	"os"
	"os/exec"

	"github.com/gogf/gf/v2/os/glog"
)

// ExecCmd 执行 gf gen 命令
// TODO 需要优化，环境变量问题
func ExecCmd(ctx context.Context) error {
	//rootDir := os.Getenv("APP_ROOT")
	//if rootDir == "" {
	//	rootDir = gfile.MainPkgPath()
	//}
	rootpath := utility.GetProjectRootSmart()

	glog.Info(ctx, "执行目录:", rootpath)

	// 创建命令
	cmd := exec.CommandContext(ctx, "gf", "gen", "ctrl")
	cmd.Dir = rootpath
	cmd.Stdout = os.Stdout // 直接将输出打印到控制台
	cmd.Stderr = os.Stderr

	// 设置环境变量
	cmd.Env = os.Environ()

	glog.Info(ctx, "执行命令:", cmd.String())

	// 执行命令
	if err := cmd.Run(); err != nil {
		glog.Error(ctx, "执行命令失败:", err)
		return err
	}

	glog.Info(ctx, "命令执行成功")
	return nil
}
