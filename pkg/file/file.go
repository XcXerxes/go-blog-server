package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}
// GetExt 获取文件后缀
func GetExt(fileName string) string  {
	return path.Ext(fileName)
}

// CheckExist 检查传入的路径是否存在
func CheckExist(src string) bool  {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission 判断是否有权限
func CheckPermission(src string) bool  {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}
// IsNotExistMkDir 如果不存在者新建文件夹
func IsNotExistMkDir(src string) error  {
	if notExist := CheckExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir 创建文件夹
func MkDir(src string) error  {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
// Open 打开文件
func Open(name string, flag int, perm os.FileMode)(*os.File, error)  {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
