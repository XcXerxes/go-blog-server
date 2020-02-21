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
	"log"
	"os"
	"strings"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

// getLogFilePath 获取日志路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// getLogFileFullPath 获取全路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	var s strings.Builder
	s.WriteString(prefixPath)
	s.WriteString(LogSaveName)
	s.WriteString(time.Now().Format(TimeFormat))
	s.WriteString(".")
	s.WriteString(LogFileExt)
	// suffixPath := s.String()
	return s.String()
}

// 打开日志文件
func openLogFile(filePath string) *os.File {
	// 返回文件信息结构描述文件，如果出现错误，会返回 *PathError
	_, err := os.Stat(filePath)
	switch {
	// 文件目录是否存在
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	// 调用文件
	if handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		log.Fatalf("fail to OpenFile：%v", err)
	}
	return handle
}

// 创建目录
func mkDir() {
	// 返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	// 创建对应的目录以及所需的子目录
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
