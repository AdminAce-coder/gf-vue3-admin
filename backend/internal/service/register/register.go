package register

import (
	"encoding/json"
	"gf-vue3-admin/utility"
	"gf-vue3-admin/utility/file"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

// 定义路由组配置结构
type RouteGroupConfig struct {
	NeedAuth bool `json:"needauth"`
	//Controller interface{} `json:"controller"`
	GroupName string `json:"groupname"`
	Enable    bool   `json:"enable"`
}

// 定义路由组
type RouteItem map[string]RouteGroupConfig

//type RouteGroup []RouteItem

var (
	ctx             = gctx.New()
	RouteGroups     = []RouteItem{}
	controllers     = make(map[string]interface{})
	routeConfigFile = filepath.Clean(gfile.Join(utility.GetProjectRootSmart(), "hack", "route_config.json"))
)

// 注册路由
func Register(name string, controller interface{}) {
	controllers[name] = controller
}

func Get(name string) interface{} {
	return controllers[name]
}

// 加载路由配置文件
func LoadRouteConfig() error {
	glog.Infof(ctx, "routeConfigFile是%s", routeConfigFile)
	// 检查路径是否存在
	_, err := os.Stat(routeConfigFile)
	// 如果错误是文件不存在
	if os.IsNotExist(err) {
		glog.Warningf(ctx, "路由配置文件不存在，创建一个新的,正在创建")
		err := file.CreateFile(routeConfigFile)
		if err != nil {
			return err
		}
	}
	// 读取文件
	glog.Infof(ctx, "正在读取文件%s", routeConfigFile)
	data, err := os.ReadFile(routeConfigFile)
	if err != nil {
		return err
		glog.Error(ctx, "读取文件失败: %w", err)
	}
	// 解析文件
	if len(data) > 0 {
		err := json.Unmarshal(data, &RouteGroups)
		if err != nil {
			glog.Error(ctx, "解析错误")
			return err
		}
	}
	return nil

}

// 新增路由
func SaveRouteConfig(newRoutes RouteItem) error {
	// 加载配置文件
	LoadRouteConfig()
	// 追加配置
	RouteGroups = append(RouteGroups, newRoutes)
	// 写入配置
	// 将配置转换为 JSON
	data, err := json.MarshalIndent(RouteGroups, "", "    ")
	if err != nil {
		return gerror.Newf("JSON编码失败: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(routeConfigFile, data, 0644); err != nil {
		return gerror.Newf("写入文件失败: %w", err)
	}

	return nil

}

// 取消注册路由
func DeleteRouteConfig(routePath string) error {
	found := false
	// 遍历查找并删除指定路由
	for _, route := range RouteGroups {
		if _, exists := route[routePath]; exists {
			delete(route, routePath)
			found = true
			break
		}
	}

	if !found {
		glog.Warningf(ctx, "删除的配置不存在: %s", routePath)
		return nil
	}

	// 重新写入配置文件
	data, err := json.MarshalIndent(RouteGroups, "", "    ")
	if err != nil {
		return gerror.Newf("JSON编码失败: %v", err)
	}

	if err := os.WriteFile(routeConfigFile, data, 0644); err != nil {
		return gerror.Newf("写入文件失败: %v", err)
	}

	return nil
}
