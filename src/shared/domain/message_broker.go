package domain

const (
	EventInsert string = "insert"
	EventUpdate string = "update"
)

type EventMessage struct {
	EventType string
	Key       string
	Message   interface{}
}
