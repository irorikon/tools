/*
 * @Author: iRorikon
 * @Date: 2023-04-17 12:03:05
 * @FilePath: \api-service\service\dingtalk\dingtalk.go
 */
package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/irorikon/api-service/config"
	"github.com/irorikon/api-service/model"
	"go.uber.org/zap"
)

type DingTalkService struct {
	AccessToken string
	Secret      string
}

func NewDingTalkService(token, secret string) *DingTalkService {
	return &DingTalkService{
		AccessToken: token,
		Secret:      secret,
	}
}

func (d *DingTalkService) URLWithTimestamp() (string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	dtu := url.URL{
		Scheme: "https",
		Host:   model.DingTalkAPI,
		Path:   "/robot/send",
	}
	value := url.Values{}
	value.Set("access_token", d.AccessToken)
	if d.Secret == "" {
		dtu.RawQuery = value.Encode()
		return dtu.String(), nil
	}
	sign, err := d.Signature(timestamp)
	if err != nil {
		dtu.RawQuery = value.Encode()
		return dtu.String(), err
	}

	value.Set("timestamp", timestamp)
	value.Set("sign", sign)
	dtu.RawQuery = value.Encode()
	config.Log.Info("URL", zap.String("URL", dtu.String()))
	return dtu.String(), nil
}

func (d *DingTalkService) Signature(timestamp string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, d.Secret)
	h := hmac.New(sha256.New, []byte(d.Secret))
	_, err := io.WriteString(h, stringToSign)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), err
}

// SendMsg
func (d *DingTalkService) SendMsg(message []byte) (*model.DingTalkResponse, error) {
	res := &model.DingTalkResponse{}

	pushURL, err := d.URLWithTimestamp()
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		pushURL,
		bytes.NewReader(message),
	)
	if err != nil {
		return res, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Charset", "utf8")

	client := new(http.Client)
	client.Timeout = time.Duration(30) * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(resByte, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *DingTalkService) FormatMsg(msg string) (*model.DingTalk, error) {
	res := &model.DingTalk{}
	err := json.Unmarshal([]byte(msg), res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
