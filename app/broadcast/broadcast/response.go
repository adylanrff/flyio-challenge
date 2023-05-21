package broadcast

import "fmt"

type Response struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}

func (r *Response) ToMap() map[string]any {
	var result = make(map[string]any)
	result["type"] = r.Type
	result["in_reply_to"] = r.InReplyTo

	return result
}

func CreateOkResponse(msgType string, msgID int) Response {
	return Response{
		Type:      fmt.Sprintf("%s_ok", msgType),
		InReplyTo: msgID,
	}
}
