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
					res, err := client.SendMsg(msgBytes)
					if err != nil {
						config.Log.Error("SendMsg error", zap.Error(err))
						continue
					}
					if res.ErrCode != 0 {
						config.Log.Error("SendMsg error", zap.Int("ErrCode", res.ErrCode), zap.String("ErrMsg", res.ErrMsg))
						continue
					}

				}
			} else {
				config.Log.Error("Message is empty")
			}
		} else {
			config.Log.Error("AccessToken or Secret is empty")
		}
	},
}

func init() {
	RootCommand.AddCommand(dingtalkCommand)
	dingtalkCommand.Flags().StringArrayVarP(&flags.Message, "message", "m", nil, "Message to send")
	dingtalkCommand.Flags().StringVarP(&flags.AccessToken, "token", "t", "2ae287d756f9cb20702b725e92ebf72a76d6741a1842a30da9f8ecf39d9d5302", "DingTalk access token")
	dingtalkCommand.Flags().StringVarP(&flags.Secret, "secret", "s", "SECadb3f65a3313d62520bd3b57fb4b739a7dcb7214e20468ac293fc7003fe9644c", "DingTalk secret")
}
