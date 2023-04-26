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

// ip search
var (
	Type     string
	Database string
	Online   bool
	IPs      []string
)

// proxy command
var (
	Address    string
	IP         string
	Username   string
	Password   string
	TCPTimeout int
	UDPTimeout int
	IPv4       bool
	IPv6       bool
)

// download command
var (
	NumConnection int
	Output        string
	Timeout       int
	Referer       string
	UserAgent     bool
	Trace         bool
	URL           []string
)
