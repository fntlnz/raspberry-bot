package sources

type Source interface {
	Name() string
	Type() string
	WaitUpdates()
	WaitFeedback()
}
