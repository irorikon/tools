/*
 * @Author: iRorikon
 * @Date: 2023-04-14 16:43:30
 * @FilePath: \api-service\main.go
 */
package main

import (
	"github.com/irorikon/tools/command"
	"github.com/irorikon/tools/config"
	"github.com/irorikon/tools/initialize"
)

func main() {
	// 初始化zap日志库
	config.Log = initialize.Zap()
	command.Execute()
}
