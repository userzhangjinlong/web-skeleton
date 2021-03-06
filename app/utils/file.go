package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//CreateIfNotExistDir 创建目录
func CreateIfNotExistDir(fileAddr string) bool {

	hasDir, err := PathExists(fileAddr)
	success := true
	if err != nil {
		fmt.Printf("PathExists(%s),err(%v)\n", fileAddr, err)
	}
	if !hasDir {
		fileErr := os.MkdirAll(fileAddr, os.ModePerm)
		if fileErr != nil {
			success = false
			logrus.Error(fileErr)
		}
	}

	return success
}

//判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//GetCurrPath 获取当前执行的绝对路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
