package apictrl

import (
	"context"
	"fmt"
	"gf-vue3-admin/internal/service/register"
	"gf-vue3-admin/utility"
	"gf-vue3-admin/utility/file"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
)

type ApiCtrl struct {
	DeleteApi      DeleteApi      // 删除API
	DeleteApiGroup DeleteApiGroup // 删除Api分组
}

// 删除APi
type DeleteApi struct {
	ApiPath  string `json:"apipath"`
	ApiGroup string `json:"apigroup"`
}
type DeleteApiGroup struct {
	ApiGroupname string `json:"apigroupname"`
	Version      string `json:"version"`
}

func (d *DeleteApi) DeleteApi(ctx context.Context) error {
	// 传入ApiPath= /api/v1/apitest/apidemota
	paths := strings.Split(d.ApiPath, "/")
	version := paths[len(paths)-3]
	apiname := paths[len(paths)-1]

	// 先删除结构体
	apigroupath := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(),
		"api", d.ApiGroup, version, d.ApiGroup+".go",
	))
	structsToRemove, err := FoudApiName(ctx, apigroupath, apiname)
	if err != nil {
		return err
	}
	noTypeLeft, err := removeStructs(apigroupath, structsToRemove)
	if err != nil {
		return err
	}

	// 如果文件中没有剩余的type声明，删除相关文件
	if noTypeLeft {
		// 构建要删除的文件路径
		interfaceFile := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(),
			"api", d.ApiGroup, fmt.Sprintf("%s.go", d.ApiGroup)))

		// 检查文件是否存在
		if gfile.Exists(interfaceFile) {
			if err := os.Remove(interfaceFile); err != nil {
				return fmt.Errorf("删除interface文件失败: %v", err)
			}
			glog.Info(ctx, "已删除interface文件:", interfaceFile)
		}
	}

	fmt.Println("apigroupath:", apigroupath)
	// 再删除控制层文件
	err = DelCtrlFile(ctx, d.ApiPath, d.ApiGroup)
	if err != nil {
		return err
	}

	// 最执行gf gen crtl
	if err = register.ExecCmd(ctx); err != nil {
		return err
	}

	return nil
}

func removeStructs(filename string, structsToRemove []string) (bool, error) {
	// 创建文件集合
	fset := token.NewFileSet()

	// 2. 解析源文件为AST（抽象语法树）
	// parser.ParseFile 将Go源代码解析成AST节点
	// nil 参数表示从文件读取内容，而不是从字符串读取
	// parser.ParseComments 表示保留注释
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

	// 创建新的声明列表，用于存储保留的声明
	var newDecls []ast.Decl
	hasTypeDecl := false // 用跟踪是否还有type声明

	// 遍历所有顶级声明
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			// 只处理类型声明（type关键字）
			if genDecl.Tok == token.TYPE {
				shouldKeep := true
				//  遍历声明中的所有规范
				for _, spec := range genDecl.Specs {
					// 检查是否是类型规范
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						// 检查是否是要删除的结构体名称
						for _, name := range structsToRemove {
							if typeSpec.Name.Name == name {
								shouldKeep = false
								break
							}
						}
					}
				}
				// 如果不是要删除的结构体，则保留
				if shouldKeep {
					hasTypeDecl = true // 如果保留了任何type声明，设置为true
					newDecls = append(newDecls, decl)
				}
			} else {
				//  非type声明直接保留
				newDecls = append(newDecls, decl)
			}
		} else {
			//  非通用声明（如函数声明）直接保留
			newDecls = append(newDecls, decl)
		}
	}

	// 更新AST
	node.Decls = newDecls

	// 创建输出文件
	output, err := os.Create(filename)
	if err != nil {
		return false, err
	}
	defer output.Close()

	// 写入修改后的代码
	if err := printer.Fprint(output, fset, node); err != nil {
		return false, err
	}

	return !hasTypeDecl, nil // 返回是否没有type声明
}

// 删除分组
func (d *DeleteApiGroup) DeleteGroup(ctx context.Context) error {

	apigroupath := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(), "api", d.ApiGroupname))
	glog.Infof(ctx, "正在删除分组,路径：%s", apigroupath)
	// 删除控制层
	apiCtrl := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(), "internal", "controller", d.ApiGroupname))
	glog.Infof(ctx, "正在删除控制层,路径：%s", apiCtrl)
	err := file.DeleteDir(apiCtrl)
	// 如果错误不是文件不存在
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// 删除API层
	err = file.DeleteDir(apigroupath)
	// 如果错误不是文件不存在
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// 取消路由注册
	routePath := fmt.Sprintf("/api/%s/%s", d.Version, d.ApiGroupname)
	glog.Infof(ctx, "正在取消注册,路径：%s", routePath)
	register.DeleteRouteConfig(routePath)

	return nil
}

// 删除控制层文件
// 传入 /api/V1/apitest
func DelCtrlFile(ctx context.Context, path, apigourp string) error {
	paths := strings.Split(path, "/")
	glog.Infof(ctx, "paths,路径：%s", paths)
	Apiname := paths[len(paths)-1]
	glog.Infof(ctx, "groupName,路径：%s", Apiname)
	ctrlPath := gfile.Join(utility.GetProjectRootSmart(), "internal", "controller", apigourp)
	glog.Infof(ctx, "ctrlPath,路径：%s", ctrlPath)
	gfile.ScanDirFunc(ctrlPath, "*.go", true, func(path string) string {
		fileName := filepath.Base(path)
		// 获取版本号后面的部分 (例如: 从 apitest_v1_api_demo_test.go 获取 api_demo_test)
		parts := strings.Split(fileName, "_")
		if len(parts) < 3 {
			return ""
		}

		afterVersion := strings.Join(parts[2:], "")
		afterVersion = strings.TrimSuffix(afterVersion, ".go")
		glog.Infof(ctx, "afterVersion,路径：%s", afterVersion)
		if strings.EqualFold(afterVersion, Apiname) {
			glog.Infof(ctx, "正在删除控制层文件,路径：%s", path)
			err := file.DeleteFile(path)
			if err != nil {
				return err.Error()
			}
		}
		return ""
	})
	return nil
}

// 找到结构体名

func FoudApiName(ctx context.Context, apigroupath, apiname string) ([]string, error) {
	content := gfile.GetContents(apigroupath)
	glog.Infof(ctx, "content: %s", content)
	// 使用 (?i) 实现不区分大小写匹配
	pattern := fmt.Sprintf(`(?i)type\s+(%s(?:req|res))\s+struct`, apiname)
	// 加载正则
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(content, -1)
	var structNames []string
	for _, match := range matches {
		if len(match) > 1 {
			// 获取实际的结构体名（保持原始大小写）
			structNames = append(structNames, match[1])
			glog.Infof(ctx, "找到结构体: %s", match[1])
		}
	}

	return structNames, nil
}
