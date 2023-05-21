package storage

type BroadcastStorage interface {
	GetSequence() (int, error)
	IncrSequence() (int, error)
	SaveMessage(msg any) error
	LoadAllMessage() ([]any, error)
}
