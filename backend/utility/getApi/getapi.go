package getApi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type RouteInfo struct {
	Path        string   `json:"path"`        // 路径
	Method      string   `json:"method"`      // 方法
	Description string   `json:"description"` // 描述
	Tags        []string `json:"tags"`        // 标签
}

type APIPaths struct {
	Paths map[string]Methods `json:"paths"`
}

type Methods struct {
	Get  *Endpoint `json:"get,omitempty"`
	Post *Endpoint `json:"post,omitempty"`
}

type Endpoint struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func GetapiInfo(ctx context.Context) *[]RouteInfo {

	// 发起 GET 请求
	resp, err := g.Client().Get(ctx, "http://127.0.0.1:5321/api/api.json")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	by := resp.Body
	// 解析 JSON 数据到结构体
	defer resp.Close()
	boby, err := io.ReadAll(by)
	if err != nil {
		g.Log().Error(ctx, err)
	}

	var apiPaths APIPaths
	err = json.Unmarshal(boby, &apiPaths)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}
	// 存储路径信息的切片
	var routes []RouteInfo

	// 遍历解析后的数据
	glog.New().Info(ctx, apiPaths.Paths)
	for path, methods := range apiPaths.Paths {
		// 构建数据
		if methods.Get != nil {
			routes = append(routes, RouteInfo{
				Path:        path,
				Method:      "GET",
				Description: methods.Get.Description,
				Tags:        methods.Get.Tags,
			})
		}
		if methods.Post != nil {
			routes = append(routes, RouteInfo{
				Path:        path,
				Method:      "POST",
				Description: methods.Post.Description,
				Tags:        methods.Post.Tags,
			})
		}
	}

	return &routes
}
