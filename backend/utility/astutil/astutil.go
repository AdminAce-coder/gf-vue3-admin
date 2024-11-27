package astutil

//
//import (

//)
//
//func AddGroupToCmd(ctx context.Context, apiGroup string) error {
//	// 获取cmd.go文件路径
//	cmdPath := filepath.Clean(gfile.Join(utility.GetProjectRootSmart(), "internal", "cmd", "cmd.go"))
//	glog.Infof(ctx, "正在修改路径%s", cmdPath)
//
//	// 创建文件集合
//	fset := token.NewFileSet()
//
//	// 解析源文件
//	node, err := parser.ParseFile(fset, cmdPath, nil, parser.ParseComments)
//	if err != nil {
//		return err
//	}
//
//	// 创建新的路由组代码
//	newGroupStr := fmt.Sprintf(`
//            // %s组，需要鉴权
//            s.Group("/api/v1/%s", func(group *ghttp.RouterGroup) {
//                group.Middleware(
//                    ghttp.MiddlewareHandlerResponse,
//                    ghttp.MiddlewareCORS,
//                    service.Middleware().Returndata, // 统一返回数据中间件
//                    service.AuthMiddleware().AuthMiddleware,
//                )
//                group.Bind(
//                    %s.NewV1(),
//                )
//            })`, apiGroup, apiGroup, apiGroup)
//
//	// 解析新的代码片段
//	newFileSet := token.NewFileSet()
//	newFile, err := parser.ParseFile(newFileSet, "", "package main\nfunc main() {"+newGroupStr+"}", parser.ParseComments)
//	if err != nil {
//		glog.Error(ctx, "解析新代码失败:", err)
//		return err
//	}
//
//	// 找到插入位置
//	var found bool
//	ast.Inspect(node, func(n ast.Node) bool {
//		if callExpr, ok := n.(*ast.CallExpr); ok {
//			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
//				if ident, ok := selExpr.X.(*ast.Ident); ok {
//					if ident.Name == "s" && selExpr.Sel.Name == "SetPort" {
//						found = true
//						glog.Info(ctx, "找到s.SetPort()位置")
//						return false
//					}
//				}
//			}
//		}
//		return true
//	})
//
//	if !found {
//		glog.Error(ctx, "未找到s.SetPort()位置")
//		return fmt.Errorf("未找到插入位置")
//	}
//
//	// 获取新的语句
//	var newStmts []ast.Stmt
//	if len(newFile.Decls) > 0 {
//		if funcDecl, ok := newFile.Decls[0].(*ast.FuncDecl); ok {
//			if len(funcDecl.Body.List) > 0 {
//				newStmts = funcDecl.Body.List
//				glog.Info(ctx, "成功获取新语句")
//			}
//		}
//	}
//
//	// 修改原AST
//	insertSuccess := false
//	ast.Inspect(node, func(n ast.Node) bool {
//		// 检查是否是 Main 变量的声明
//		if genDecl, ok := n.(*ast.GenDecl); ok {
//			for _, spec := range genDecl.Specs {
//				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
//					for _, name := range valueSpec.Names {
//						if name.Name == "Main" {
//							// 找到 Main 变量的声明
//							if compositeLit, ok := valueSpec.Values[0].(*ast.CompositeLit); ok {
//								for _, elt := range compositeLit.Elts {
//									if kvExpr, ok := elt.(*ast.KeyValueExpr); ok {
//										if ident, ok := kvExpr.Key.(*ast.Ident); ok {
//											if ident.Name == "Func" {
//												// 在 Func 字段中查找 s.SetPort()
//												if funcLit, ok := kvExpr.Value.(*ast.FuncLit); ok {
//													blockStmt := funcLit.Body
//													// 找到插入位置
//													for i, stmt := range blockStmt.List {
//														if exprStmt, ok := stmt.(*ast.ExprStmt); ok {
//															if callExpr, ok := exprStmt.X.(*ast.CallExpr); ok {
//																if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
//																	if ident, ok := selExpr.X.(*ast.Ident); ok {
//																		if ident.Name == "s" && selExpr.Sel.Name == "SetPort" {
//																			// 在s.SetPort()之前插入新的语句
//																			newList := make([]ast.Stmt, 0)
//																			newList = append(newList, blockStmt.List[:i]...)
//																			newList = append(newList, newStmts...)
//																			newList = append(newList, blockStmt.List[i:]...)
//																			blockStmt.List = newList
//																			insertSuccess = true
//																			glog.Info(ctx, "成功插入新代码")
//																			return false
//																		}
//																	}
//																}
//															}
//														}
//													}
//												}
//											}
//										}
//									}
//								}
//							}
//						}
//					}
//				}
//			}
//		}
//		return true
//	})
//
//	if !insertSuccess {
//		glog.Error(ctx, "插入代码失败")
//		return fmt.Errorf("插入代码失败")
//	}
//
//	// 添加必要的导入
//	if !astutil.AddImport(fset, node, fmt.Sprintf("gf-vue3-admin/internal/controller/%s", apiGroup)) {
//		glog.Error(ctx, "添加导入失败")
//		return fmt.Errorf("添加导入失败")
//	}
//
//	// 创建输出文件
//	output, err := os.Create(cmdPath)
//	if err != nil {
//		glog.Error(ctx, "创建输出文件失败:", err)
//		return err
//	}
//	defer output.Close()
//
//	// 设置格式化选项
//	config := printer.Config{
//		Mode:     printer.UseSpaces | printer.TabIndent,
//		Tabwidth: 4,
//	}
//
//	// 写入修改后的代码
//	if err := config.Fprint(output, fset, node); err != nil {
//		glog.Error(ctx, "写入文件失败:", err)
//		return err
//	}
//
//	glog.Info(ctx, "成功修改cmd.go文件")
//	return nil
//}
