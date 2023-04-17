/*
 * @Author: iRorikon
 * @Date: 2023-04-17 14:44:34
 * @FilePath: \api-service\command\flags\flags.go
 */
package flags

// root command
var (
	Debug   bool
	LogPath string
)

// dingtalk command
var (
	AccessToken string
	Secret      string
	Message     []string
)
