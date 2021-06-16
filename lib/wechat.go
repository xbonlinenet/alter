package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
)

const (
	SendMessageUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="
)

var (
	httpClient *http.Client
)

func init() {
	timeout := time.Duration(30 * time.Second)
	httpClient = &http.Client{
		Timeout: timeout,
	}

}

// WechatChannel 发送微信通知的客户端
type WechatChannel struct {
	CorpID        string
	CorpSecret    string
	AgentID       int
	tokenExpireAt int64
	accessToken   string
}

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}

type CustomServiceMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	AgentID int            `json:"agentid"`
	Text    TextMsgContent `json:"text"`
}

func (client *WechatChannel) getAccessTokenURL() string {
	return fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", client.CorpID, client.CorpSecret)
}

func (client *WechatChannel) fetchAccessToken() {
	now := time.Now().Unix()
	if now >= client.tokenExpireAt {
		url := client.getAccessTokenURL()
		resp, err := httpClient.Get(url)

		if err != nil || resp.StatusCode != http.StatusOK {
			glog.Errorf("get access_token error")
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if bytes.Contains(body, []byte("access_token")) {
			atr := AccessTokenResponse{}
			err = json.Unmarshal(body, &atr)
			if err != nil {
				panic(err)
			}
			client.accessToken = atr.AccessToken
			client.tokenExpireAt = time.Now().Unix() + int64(atr.ExpiresIn) - 60
		} else {
			glog.Errorf("get access_token error: %s", string(body))
			return
		}
	}
}

// SendMessage 发送消息给制定用户列表
func (client *WechatChannel) SendMessage(users []string, robotUrls []string, msg string) error {
	client.fetchAccessToken()

	toUser := strings.Join(users, "|")
	csMsg := &CustomServiceMsg{
		ToUser:  toUser,
		MsgType: "text",
		AgentID: client.AgentID,
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(csMsg, " ", "  ")
	if err != nil {
		return err
	}

	resp, err := httpClient.Post(strings.Join([]string{SendMessageUrl, client.accessToken}, ""), "application/json; encoding=utf-8", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	resp.Body.Close()
	// 发送到机器人
	if len(robotUrls) > 0 {
		for _, url := range robotUrls {
			client.SendGroup(url, msg)
		}
	}
	return nil
}

// SendGroup 通过企业微信机器人 发送消息
func (client *WechatChannel) SendGroup(url string, msg string) error {
	type Payload struct {
		MSGType string         `json:"msgtype"`
		Text    TextMsgContent `json:"text"`
	}
	payload := &Payload{
		MSGType: "text",
		Text:    TextMsgContent{Content: msg},
	}
	body, err := json.MarshalIndent(payload, "", "")

	resp, err := httpClient.Post(url, "application/json; encoding=utf-8", bytes.NewReader(body))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
