package uuid

type Response struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
	ID        string `json:"id"`
}

func (r *Response) ToMap() map[string]any {
	var result = make(map[string]any)
	result["type"] = r.Type
	result["id"] = r.ID

	return result
}

func CreateOKResponse(id string, msgID int) Response {
	return Response{
		ID:        id,
		InReplyTo: msgID,
		Type:      "generate_ok",
	}
}
