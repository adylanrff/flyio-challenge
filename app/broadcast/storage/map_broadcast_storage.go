package storage

import (
	"errors"
	"fmt"
	"sync"
)

const (
	StorageSequenceKey   = "sequence"
	StorageMessageFormat = "message-%d"
)

type MapMessageStorage struct {
	sync.Map
}

func (m *MapMessageStorage) GetSequence() (int, error) {
	value, _ := m.Map.LoadOrStore(StorageSequenceKey, 0)
	castedValue, ok := value.(int)
	if !ok {
		return 0, errors.New("invalid value")
	}

	return castedValue, nil
}
func (m *MapMessageStorage) IncrSequence() (int, error) {
	curSequence, err := m.GetSequence()
	if err != nil {
		return -1, err
	}

	m.Map.Store(StorageSequenceKey, curSequence+1)
	return m.GetSequence()
}

func (m *MapMessageStorage) SaveMessage(msg any) error {
	curSequence, err := m.IncrSequence()
	if err != nil {
		return err
	}

	m.Map.Store(fmt.Sprintf(StorageMessageFormat, curSequence), msg)
	return nil
}

func (m *MapMessageStorage) LoadAllMessage() ([]any, error) {
	sequence, err := m.GetSequence()
	if err != nil {
		// heuristic, make it 10 first
		sequence = 10
	}

	var result = make([]any, sequence)

	m.Map.Range(func(key, value any) bool {
		result = append(result, value)
		return true
	})

	return result, nil
}

var globalMessageStorage *MapMessageStorage

func InitMessageStorage() {
	globalMessageStorage = &MapMessageStorage{}
}

func GetMessageStore() *MapMessageStorage {
	return globalMessageStorage
}
