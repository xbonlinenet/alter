package lib

type Channel interface {
	SendMessage(users []string, msg string) error
}
