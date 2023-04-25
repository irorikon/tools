package command

import (
	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/config"
	"github.com/spf13/cobra"
	"github.com/txthinking/socks5"
	"go.uber.org/zap"
)

var socksCommand = &cobra.Command{
	Use:   "socks",
	Short: "socks Server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := socks5.NewClassicServer(
			flags.Address,
			flags.IP,
			flags.Username,
			flags.Password,
			flags.TCPTimeout,
			flags.UDPTimeout,
		)
		if err != nil {
			config.Log.Error("Error", zap.String("socks error", err.Error()))
		}
		config.Log.Info("Start socks server", zap.String("address", flags.Address))
		config.Log.Error("Error", zap.String("socks error", server.ListenAndServe(server.Handle).Error()))
	},
}

func init() {
	RootCommand.AddCommand(socksCommand)
	socksCommand.Flags().StringVarP(&flags.Address, "address", "a", ":1080", "Address")
	socksCommand.Flags().StringVar(&flags.IP, "ip", "", "IP")
	socksCommand.Flags().StringVarP(&flags.Username, "username", "u", "", "Username")
	socksCommand.Flags().StringVarP(&flags.Password, "password", "p", "", "Password")
	socksCommand.Flags().IntVar(&flags.TCPTimeout, "tcp-timeout", 300, "TCP Timeout")
	socksCommand.Flags().IntVar(&flags.UDPTimeout, "udp-timeout", 300, "UDP Timeout")
	socksCommand.Flags().BoolVarP(&flags.IPv4, "ipv4", "4", true, "IPv4")
	socksCommand.Flags().BoolVarP(&flags.IPv6, "ipv6", "6", false, "IPv6")
}
