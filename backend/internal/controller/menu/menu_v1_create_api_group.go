package menu

import (
	"context"
	"gf-vue3-admin/api/menu/v1"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) CreateApiGroup(ctx context.Context, req *v1.CreateApiGroupReq) (res *v1.CreateApiGroupRes, err error) {
	// 判断apiFile是否存在,不存在则创建
	apiFilePath := gfile.Join(gfile.Pwd(), "api", req.ApiGroupName, req.Version, req.ApiGroupName+".go")
	if !gfile.Exists(apiFilePath) {
		glog.New().Error(ctx, "api file does not exist")
		// 创建api文件
		if err = CreateGroupFile(apiFilePath, req.Version); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

// / 创建初始文件
func CreateGroupFile(pathname, version string) error {
	file, err := gfile.Create(pathname)
	if err != nil {
		return err
	}
	defer file.Close()
	// 追加写入package
	_, err = file.WriteString("package " + version)
	if err != nil {
		return err
	}
	return nil
}
