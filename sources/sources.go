package sources

type Source interface {
	Name() string
	Type() string
	WaitUpdates()
	WaitFeedback()
}

var updates = make(chan *Message)
var feedback = make(chan *Message)

func Updates() chan *Message {
	return updates
}

func Feedback() chan *Message {
	return feedback
}
