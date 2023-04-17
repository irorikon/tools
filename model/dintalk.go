/*
 * @Author: iRorikon
 * @Date: 2023-04-14 17:29:52
 * @FilePath: \api-service\model\dintalk.go
 */
package model

// https://oapi.dingtalk.com/robot/send?access_token=xxx
const DingTalkAPI string = "oapi.dingtalk.com"

type DingTalk struct {
	AT         ATUser        `json:"at"`
	Text       TextMsg       `json:"text"`
	Link       LinkMsg       `json:"link"`
	ActionCard ActionCardMsg `json:"actionCard"`
	FeedCard   FeedCardMsg   `json:"feedCard"`
	Markdown   MarkdownMsg   `json:"markdown"`
	MsgType    string        `json:"msgtype"`
}

type ATUser struct {
	AtMobiles []string `json:"atMobiles"`
	AtUserIds []string `json:"atUserIds"`
	IsAtAll   bool     `json:"isAtAll"`
}

type TextMsg struct {
	Content string `json:"content"`
}

type LinkMsg struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

type ActionCardMsg struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []Btn  `json:"btns"`
}

type Btn struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

type FeedCardMsg struct {
	Links []Link `json:"links"`
}

type Link struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}

type MarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type DingTalkResponse struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
}
