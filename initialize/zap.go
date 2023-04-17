/*
 * @Author: iRorikon
 * @Date: 2023-04-17 15:16:00
 * @FilePath: \api-service\initialize\zap.go
 */
package initialize

import (
	"fmt"
	"os"

	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/config"
	"github.com/irorikon/api-service/initialize/internal"
	"github.com/irorikon/api-service/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok := util.DirExist(flags.LogPath); !ok {
		fmt.Printf("create %v directory\n", flags.LogPath)
		_ = os.Mkdir(flags.LogPath, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if config.ZapShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
