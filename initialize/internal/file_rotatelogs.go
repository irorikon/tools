/*
 * @Author: iRorikon
 * @Date: 2023-04-17 15:42:31
 * @FilePath: \api-service\initialize\internal\file_rotatelogs.go
 */
package internal

import (
	"os"
	"path"
	"time"

	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

type fileRotateLogs struct{}

var FileRotateLogs = new(fileRotateLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(flags.LogPath, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(config.ZapMaxAge)*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if config.ZapLogToConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
