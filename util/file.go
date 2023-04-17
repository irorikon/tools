/*
 * @Author: iRorikon
 * @Date: 2023-04-17 15:19:11
 * @FilePath: \api-service\util\file.go
 */
package util

import "os"

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// DirExist 判断目录是否存在
func DirExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return fi.IsDir()
	}
	return !os.IsNotExist(err)
}
