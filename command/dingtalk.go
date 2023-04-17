/*
 * @Author: iRorikon
 * @Date: 2023-04-17 14:28:29
 * @FilePath: \api-service\command\dingtalk.go
 */
package command

import (
	"encoding/json"
	"fmt"

	"github.com/irorikon/api-service/command/flags"
	"github.com/irorikon/api-service/config"
	"github.com/irorikon/api-service/model"
	"github.com/irorikon/api-service/service/dingtalk"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var dingtalkCommand = &cobra.Command{
	Use:   "dingtalk",
	Short: "Send message to DingTalk",
	Run: func(cmd *cobra.Command, args []string) {
		if flags.AccessToken != "" && flags.Secret != "" {
			config.Log.Info("DingTalk access token and secret is set", zap.String("AccessToken", flags.AccessToken), zap.String("Secret", flags.Secret))
			client := dingtalk.NewDingTalkService(flags.AccessToken, flags.Secret)
			if flags.Message != nil {
				for _, msgString := range flags.Message {
					fmt.Println(msgString)
					message := new(model.DingTalk)
					err := json.Unmarshal([]byte(msgString), &message)
					if err != nil {
						config.Log.Error("Unmarshal error", zap.Error(err))
						continue
					}
					msgBytes, err := json.Marshal(message)
					if err != nil {
						config.Log.Error("Marshal error", zap.Error(err))
						continue
					}
					if res, err := client.SendMsg(msgBytes); err != nil {
						config.Log.Error("SendMsg error", zap.Error(err))
						continue
					} else {
						if res.ErrCode != 0 {
							config.Log.Error("SendMsg error", zap.Int("ErrCode", res.ErrCode), zap.String("ErrMsg", res.ErrMsg))
							continue
						}
					}
				}
			} else {
				fmt.Println("Message is empty")
			}
		}
	},
}

func init() {
	RootCommand.AddCommand(dingtalkCommand)
	dingtalkCommand.Flags().StringArrayVarP(&flags.Message, "message", "m", nil, "Message to send")
	dingtalkCommand.Flags().StringVarP(&flags.AccessToken, "token", "t", "", "DingTalk access token")
	dingtalkCommand.Flags().StringVarP(&flags.Secret, "secret", "s", "", "DingTalk secret")
}
