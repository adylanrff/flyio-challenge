package broadcast

import (
	"encoding/json"
	"errors"
)

type Request struct {
	Type    string `json:"type"`
	MsgID   int    `json:"msg_id"`
	Message any    `json:"message"`
}

func (r *Request) Validate() error {
	if r.Type != "broadcast" {
		return errors.New("invalid param")
	}

	return nil
}

func ParseRequest(msgBody json.RawMessage) (Request, error) {
	var parsedRequest Request
	if err := json.Unmarshal(msgBody, &parsedRequest); err != nil {
		return Request{}, err
	}

	return parsedRequest, nil
}
