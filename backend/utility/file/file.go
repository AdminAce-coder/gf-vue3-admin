package file

import "github.com/gogf/gf/v2/os/gfile"

// 删除文件
func DeleteFile(path string) error {

	err := gfile.RemoveFile(path)
	if err != nil {
		return err
	}
	return nil
}

// 删除文件夹
func DeleteDir(path string) error {
	err := gfile.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// 创建文件夹
func CreateDir(path string) error {
	err := gfile.Mkdir(path)
	if err != nil {
		return err
	}
	return nil
}

// 创建文件
func CreateFile(path string) error {
	_, err := gfile.Create(path)
	if err != nil {
		return err
	}
	return nil
}
