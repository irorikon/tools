/*
 * @Author: iRorikon
 * @Date: 2023-04-26 11:04:29
 * @FilePath: \api-service\command\download.go
 */
package command

import (
	"context"
	"fmt"
	"os"

	"github.com/irorikon/tools/command/flags"
	"github.com/irorikon/tools/service/download"
	"github.com/irorikon/tools/util"
	"github.com/spf13/cobra"
)

var downloadCommand = &cobra.Command{
	Use:   "get",
	Short: "download files online",
	Run: func(cmd *cobra.Command, args []string) {
		client := download.New()
		if err := client.Run(context.Background(), util.Version, os.Args[1:]); err != nil {
			if client.Trace {
				fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
			} else {
				fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
			}
			os.Exit(1)
		}
	},
}

func init() {
	RootCommand.AddCommand(downloadCommand)
	downloadCommand.Flags().IntVarP(&flags.NumConnection, "procs", "p", 4, "the number of connections for a single URL")
	downloadCommand.Flags().StringVarP(&flags.Output, "output", "o", "", "output file to <filename>")
	downloadCommand.Flags().BoolVar(&flags.UserAgent, "user-agent", false, "random user-agent")
	downloadCommand.Flags().StringVarP(&flags.Referer, "referer", "r", "", "identify as <referer>")
	downloadCommand.Flags().IntVarP(&flags.Timeout, "timeout", "t", 10, "timeout of checking request in seconds")
	downloadCommand.Flags().StringArrayVarP(&flags.URL, "url", "u", nil, "url to download")
	downloadCommand.Flags().BoolVar(&flags.Trace, "trace", false, "display detail error messages")
}
