package read

import (
	"fmt"
)

type Response struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
	Messages  []any  `json:"messages"`
}

func (r *Response) ToMap() map[string]any {
	var result = make(map[string]any)
	result["type"] = r.Type
	result["in_reply_to"] = r.InReplyTo
	result["messages"] = r.Messages

	return result
}

func CreateOkResponse(msgType string, msgID int, messages []any) Response {
	return Response{
		Type:      fmt.Sprintf("%s_ok", msgType),
		InReplyTo: msgID,
		Messages:  messages,
	}
}
