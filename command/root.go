/*
 * @Author: iRorikon
 * @Date: 2023-04-14 16:44:49
 * @FilePath: \api-service\command\root.go
 */
package command

import (
	"fmt"
	"os"

	"github.com/irorikon/tools/command/flags"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "api-service",
	Short: "API service program that integrates multiple functions.",
	Long:  "API service program that integrates multiple functions.",
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCommand.PersistentFlags().BoolVarP(&flags.Debug, "debug", "d", false, "Start with debug mode")
	RootCommand.PersistentFlags().StringVarP(&flags.LogPath, "log-path", "l", "logs", "Log path")
}
