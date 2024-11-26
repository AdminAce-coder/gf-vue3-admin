package astutil

import (
	"context"
	"gf-vue3-admin/utility"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gogf/gf/v2/os/gfile"
)

func TestCreateApi(t *testing.T) {
	ctx := context.Background()

	// 测试前备份文件
	cmdPath := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(), "internal", "cmd", "cmd.go"))
	backupPath := cmdPath + ".backup"
	if err := gfile.Copy(cmdPath, backupPath); err != nil {
		t.Fatal("备份文件失败:", err)
	}
	defer gfile.Copy(backupPath, cmdPath) // 测试后恢复文件

	if err := AddGroupToCmd(ctx, "apitest"); err != nil {
		t.Error(err)
	}

	// 验证文件是否被修改
	content := gfile.GetContents(cmdPath)

	if !strings.Contains(content, "apitest.NewV1()") {
		t.Error("文件未被正确修改")
	}
}
