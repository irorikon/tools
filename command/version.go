/*
 * @Author: iRorikon
 * @Date: 2023-04-14 17:07:11
 * @FilePath: \api-service\command\version.go
 */
package command

import (
	"fmt"
	"os"

	"github.com/irorikon/api-service/util"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show current version of api-service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(util.VersionTpl())
		os.Exit(0)
	},
}

func init() {
	RootCommand.AddCommand(versionCommand)
}
