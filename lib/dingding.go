package lib

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	DingDingUrl = "http://xbopalarm.appsflyer.club/api/v1"
)

type DingDingChannel struct {
}

func (d *DingDingChannel) SendMessage(users []string, msg string) error {

	title := fmt.Sprintf("To:%s", strings.Join(users, ", "))
	url := fmt.Sprintf("%s?party=%d&title=%s&content=%s&angentid=%d", DingDingUrl, 5, title, url.QueryEscape(msg), 1000009)

	fmt.Printf("url: %s\n", url)
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
