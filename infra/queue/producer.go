package queue

type Producer interface {
	SendMessage(msg []byte) error
}
