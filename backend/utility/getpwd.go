package utility

import (
	"os"
	"path/filepath"
	"runtime"
)

// 获取根路径
func GetProjectRootSmart() string {
	// 1. 首先尝试环境变量
	if root := os.Getenv("PROJECT_ROOT"); root != "" {
		if validateProjectRoot(root) {
			return root
		}
	}

	// 2. 尝试运行时路径
	if root := GetProjectRootByRuntime(); root != "" {
		if validateProjectRoot(root) {
			return root
		}
	}

	// 3. 尝试可执行文件路径
	if execRoot := GetProjectRoot(); execRoot != "" {
		if validateProjectRoot(execRoot) {
			return execRoot
		}
	}

	// 4. 最后尝试当前工作目录
	workDir, err := os.Getwd()
	if err == nil && validateProjectRoot(workDir) {
		return workDir
	}

	return ""
}

// 根据运行时获取
func GetProjectRootByRuntime() string {
	// 获取当前文件的绝对路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	// 回溯到项目根目录
	return filepath.Dir(filepath.Dir(filename))
}

// 验证目录是否为有效的项目根目录
func validateProjectRoot(path string) bool {
	configDir := filepath.Join(path, "internal", "cmd")
	if _, err := os.Stat(configDir); err == nil {
		return true
	}
	return false
}

// 使用可执行文件位置作为基准
func GetProjectRoot() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	// 获取可执行文件所在目录
	return filepath.Dir(execPath)
}
