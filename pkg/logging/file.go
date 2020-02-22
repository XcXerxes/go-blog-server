/*
 * @Description: 日志
 * @Author: leo
 * @Date: 2020-02-21 19:09:05
 * @LastEditors: leo
 * @LastEditTime: 2020-02-21 19:57:42
 */

package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/XcXerxes/go-blog-server/pkg/file"
	"github.com/XcXerxes/go-blog-server/pkg/setting"
)

// getLogFilePath 获取日志路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", setting.AppSetting.LogSavePath)
}

// getLogFileName 获取全路径
func getLogFileName() string {
	prefixPath := getLogFilePath()
	appSetting := setting.AppSetting
	var s strings.Builder
	s.WriteString(prefixPath)
	s.WriteString(appSetting.LogSaveName)
	s.WriteString(time.Now().Format(appSetting.TimeFormat))
	s.WriteString(".")
	s.WriteString(appSetting.LogFileExt)
	// suffixPath := s.String()
	return s.String()
}

// 打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	// 返回与当前目录对应的根路径名
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}
	src := dir + "/" + filePath
	// 是否有权限
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}
	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return f, nil
}
