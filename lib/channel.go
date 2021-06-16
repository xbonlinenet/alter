package lib

type Channel interface {
	SendMessage(users []string, robotUrls []string, msg string) error
}
