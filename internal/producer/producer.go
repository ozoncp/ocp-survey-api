package producer

type EventType int

const (
	Undefined EventType = iota
	Create
	Update
	Delete
)

var eventTypeStr = []string{
	"Undefined",
	"Create",
	"Update",
	"Delete",
}

type Event interface {
	Value() string
}

type Producer interface {
	Send(topic string, event Event) error
	Close()
}

func (e EventType) String() string {
	if i := int(e); i < len(eventTypeStr) {
		return eventTypeStr[i]
	}
	return ""
}
