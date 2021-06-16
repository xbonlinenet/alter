package lib

import (
	"log"
	"net"

	"github.com/go-redis/redis"
)

// Client 微信报警客户端
type Client struct {
	client    redis.Cmdable
	server    string
	users     []string
	robotUrls []string
	host      string
}

// NewClient 创建新的客户端
func NewClient(client redis.Cmdable, users, robotUrls []string, server string) (*Client, error) {

	host := GetOutboundIP().String()
	return &Client{
		client:    client,
		server:    server,
		users:     users,
		robotUrls: robotUrls,
		host:      host,
	}, nil
}

// AlterUsers 发送报警信息到统一报警平台
func (c *Client) AlterUsers(users []string, message string, detail string, errorID string) error {
	errMessage := ErrorMessage{
		Host:      c.host,
		Server:    c.server,
		Users:     users,
		RobotUrls: c.robotUrls,
		ErrorID:   errorID,
		Message:   message,
		Detail:    detail,
	}

	strMessage, err := EncodeErrorMessage(errMessage)
	if err != nil {
		return err
	}
	_, err = c.client.LPush(RedisErrListKey, strMessage).Result()
	return err
}

// Alter 发送报警信息到统一报警平台
func (c *Client) Alter(message string, detail string, errorID string) error {
	return c.AlterUsers(c.users, message, detail, errorID)
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
