/*
 * @Author: iRorikon
 * @Date: 2023-04-14 17:08:58
 * @FilePath: \api-service\pkg\version\version.go
 */
package version

import (
	"fmt"
	"runtime"
)

var (
	version   string
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
	return fmt.Sprintf(tpl, version, commit, buildDate, goVersion, runtime.GOOS+"/"+runtime.GOARCH)
}
