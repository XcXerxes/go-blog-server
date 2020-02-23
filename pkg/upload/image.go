/*
 * @Description: 上传图片逻辑
 * @Author: leo
 * @Date: 2020-02-23 14:03:08
 * @LastEditors: leo
 * @LastEditTime: 2020-02-23 20:45:17
 */
package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/XcXerxes/go-blog-server/pkg/file"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
	"github.com/XcXerxes/go-blog-server/pkg/util"
)

// GetImageFullUrl 获取图片的全路径
func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

// GetImagePath 获取图片的保存路径
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

// GetImageName 获取图片的名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	// 返回 name 中 不包含 ext后缀的内容
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMd5(fileName)
	return fileName + ext
}

// GetImageFullPath 图片的全
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

// CheckImageExt 检查上传的图片文件 是否符合 支持 .png|.jpg|.jpeg
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImageSize 检查图片的大小
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}

// CheckImage 检查图片
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	return nil
}
