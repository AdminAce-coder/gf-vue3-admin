package menu

import (
	"context"
	"fmt"
	v1 "gf-vue3-admin/api/menu/v1"
	"gf-vue3-admin/internal/service/register"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

type Apifile struct {
	version     string
	apiFile     string
	apiPathName string
}

func (c *ControllerV1) CreateApi(ctx context.Context, req *v1.CreateApiReq) (res *v1.CreateApiRes, err error) {

	//}
	apiGroup := strings.Trim(req.ApiGroup, "/\\")
	apiVersion := strings.Trim(req.ApiVersion, "/\\")
	apigroupath := filepath.Clean(gfile.Join(gfile.Pwd(), "api", apiGroup, apiVersion, apiGroup+".go"))
	glog.New().Infof(ctx, "处理前的分组: %s", req.ApiGroup)
	glog.New().Infof(ctx, "处理后的分组: %s", apiGroup)
	glog.New().Infof(ctx, "最终路径是: %s", apigroupath)
	// 生成API结构体
	if err = CreateApiStruct(apigroupath, req); err != nil {
		return nil, err
	}

	// 生成控制层
	if err = register.ExecCmd(ctx); err != nil {
		return nil, err
	}
	return &v1.CreateApiRes{}, nil
}

// 生成api 结构体
func CreateApiStruct(apiFilePath string, req *v1.CreateApiReq) error {
	// apiPathName 全部转为小写
	apipath := strings.ToLower(req.ApiName)
	// 打开文件
	file, err := gfile.OpenWithFlag(apiFilePath, os.O_RDWR|os.O_APPEND)
	if err != nil {
		return err
	}
	defer file.Close()

	// 构建结构体内容
	var structContent strings.Builder

	// 添加换行
	structContent.WriteString("\n\n")

	// 构建请求结构体
	structContent.WriteString(fmt.Sprintf("type %sReq struct {\n", req.ApiName))
	structContent.WriteString(fmt.Sprintf("\tg.Meta `path:\"/%s\" method:\"%s\" tags:\"%s\" dc:\"%s\"`\n",
		apipath, req.Method, req.ApiGroup, req.Description))

	// 如果是POST请求，添加参数字段
	if req.Method == "post" && len(req.Parameters) > 0 {
		glog.New().Info(gctx.New(), "正在进行参数写入...")
		for _, param := range req.Parameters {
			// 添加字段注释
			if param.Description != "" {
				structContent.WriteString(fmt.Sprintf("\t// %s\n", param.Description))
			}
			// 根据是否必须添加 binding 标签
			if param.Required {
				structContent.WriteString(fmt.Sprintf("\t%s %s `v:\"required\" dc:\"%s\"`\n",
					param.ParameterName, param.DataType, param.Description))
			} else {
				structContent.WriteString(fmt.Sprintf("\t%s %s `dc:\"%s\"`\n",
					param.ParameterName, param.DataType, param.Description))
			}
		}
	}
	structContent.WriteString("}\n\n")

	// 构建响应结构体
	structContent.WriteString(fmt.Sprintf("type %sRes struct {\n", req.ApiName))
	structContent.WriteString("}\n")

	// 写入文件
	_, err = file.WriteString(structContent.String())
	return err
}
