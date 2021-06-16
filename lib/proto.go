package lib

import "encoding/json"

// ErrorMessage 错误信息
type ErrorMessage struct {
	Host           string   `json:"host"`
	Server         string   `json:"server"`
	Users          []string `json:"users"`
	RobotUrls      []string `json:"robot_urls"`
	AlterThreshold int      `json:"alter_threshold"`
	ErrorID        string   `json:"error_id"`
	Message        string   `json:"message"`
	Detail         string   `json:"detail"`
}

// DecodeErrorMessage 反序列化
func DecodeErrorMessage(data []byte) (ErrorMessage, error) {
	var message ErrorMessage
	err := json.Unmarshal(data, &message)
	return message, err
}

// EncodeErrorMessage 序列化成字符串
func EncodeErrorMessage(message ErrorMessage) (string, error) {
	data, err := json.Marshal(&message)
	if err != nil {
		return "", err
	}
	return string(data), err
}
