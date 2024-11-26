package menu

import (
	"context"
	v1 "gf-vue3-admin/api/menu/v1"
	"gf-vue3-admin/internal/service/register"
	"gf-vue3-admin/utility"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) CreateApiGroup(ctx context.Context, req *v1.CreateApiGroupReq) (res *v1.CreateApiGroupRes, err error) {
	// 判断apiFile是否存在,不存在则创建 /api/v1/auth
	path := req.ApiPath // /api/v1/apitest
	//分割
	split := strings.Split(path, "/")
	apiFilePath := gfile.Join(utility.GetProjectRootSmart(), "api", split[len(split)-1], split[len(split)-2], split[len(split)-1]+".go")
	if !gfile.Exists(apiFilePath) {
		glog.New().Error(ctx, "api file does not exist")
		// 创建api文件
		if err = CreateGroupFile(apiFilePath, split[len(split)-2]); err != nil {
			return nil, err
		}
	}
	re := register.RouteItem{
		path: req.Register,
	}
	err = register.SaveRouteConfig(re)
	if err != nil {
		return nil, err
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

// 写入路由表
