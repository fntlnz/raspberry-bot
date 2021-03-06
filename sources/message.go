package sources

type Message struct {
	SourceName string
	Sender     interface{}
	Body       interface{}
}

var updates = make(chan *Message)
var feedback = make(chan *Message)

func Updates() chan *Message {
	return updates
}

func Feedback() chan *Message {
	return feedback
}
