package topology

import (
	"encoding/json"
	"errors"
)

type Request struct {
	Type string `json:"type"`
}

func (r *Request) ToMap() map[string]any {
	return map[string]any{
		"type": r.Type,
	}
}

func (r *Request) Validate() error {
	if r.Type != "topology" {
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
