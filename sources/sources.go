package sources

type Source interface {
	SourceName() string
	WaitUpdates()
	WaitFeedback()
	Updates() <-chan *Message
	Feedback() chan<- *Message
}
