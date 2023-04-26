/*
 * @Author: iRorikon
 * @Date: 2023-04-14 17:08:58
 * @FilePath: \api-service\util\version.go
 */
package util

import (
	"fmt"
	"runtime"
)

var (
	Version   string
	buildDate string
	commit    string
	goVersion string
	tpl       string = `api-service, version %s (branch: HEAD, revision: %s)
  build user:       irorikon@88.com
  build date:       %s
  go version:       %s
  platform:         %s`
)

func VersionTpl() string {
	return fmt.Sprintf(tpl, Version, commit, buildDate, goVersion, runtime.GOOS+"/"+runtime.GOARCH)
}
