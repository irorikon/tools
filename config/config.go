/*
 * @Author: iRorikon
 * @Date: 2023-04-17 16:01:15
 * @FilePath: \api-service\config\config.go
 */
package config

import (
	"github.com/irorikon/tools/command/flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap Config
var (
	Log *zap.Logger
)

const (
	ZapMaxAge        int    = 72
	ZapFormat        string = "console"
	ZapPrefix        string = "[api-service]"
	ZapStacktraceKey string = "stacktrace"
	ZapLevel         string = "info"
	ZapShowLine      bool   = true
	ZapLogToConsole  bool   = true
)

func TransportLevel() zapcore.Level {
	var level string
	if flags.Debug {
		level = "debug"
	} else {
		level = ZapLevel
	}
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// IP Search Config
