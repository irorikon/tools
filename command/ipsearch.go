package command

import (
	"encoding/json"
	"fmt"

	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/config"
	"github.com/irorikon/api-service/service/ipsearch"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var ipsearchCommand = &cobra.Command{
	Use:   "ipsearch",
	Short: "Query the details of an IP address",
	Run: func(cmd *cobra.Command, args []string) {
		if flags.IPs == nil {
			config.Log.Error("No IP address given")
		}
		search := ipsearch.NewIPSearch(flags.Type, flags.Database, flags.Online, flags.IPs)
		if flags.Database != "" && flags.Type != "" {
			ipinfo, err := search.SearchByDatabase()
			if err != nil {
				config.Log.Error("Error", zap.String("search error", err.Error()))
			}
			formatJson, err := json.Marshal(ipinfo)
			if err != nil {
				config.Log.Error("Error", zap.String("format error", err.Error()))
			}
			fmt.Println(string(formatJson))
		} else if flags.Online {
			ipinfo, err := search.SearchByOnline()
			if err != nil {
				config.Log.Error("Error", zap.String("search error", err.Error()))
			}
			formatJson, err := json.Marshal(ipinfo)
			if err != nil {
				config.Log.Error("Error", zap.String("format error", err.Error()))
			}
			fmt.Println(string(formatJson))
		} else {
			config.Log.Error("No database or online specified.")
		}
	},
}

func init() {
	RootCommand.AddCommand(ipsearchCommand)
	ipsearchCommand.Flags().StringArrayVarP(&flags.IPs, "ip", "i", nil, "Select IP or IP List.")
	ipsearchCommand.Flags().StringVarP(&flags.Type, "type", "t", "ip2location", "Database Type(default ip2location)")
	ipsearchCommand.Flags().StringVarP(&flags.Database, "file", "f", "", "Specified ip database file.")
	ipsearchCommand.Flags().BoolVarP(&flags.Online, "online", "o", false, "Search Online.")
}
