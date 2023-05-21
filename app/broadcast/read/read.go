package read

import "flyio-challenge/app/broadcast/storage"

func Read() ([]any, error) {
	messageStore := storage.GetMessageStore()
	return messageStore.LoadAllMessage()
}
