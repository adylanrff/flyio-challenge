package broadcast

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func Handler(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		request, err := ParseRequest(msg.Body)
		if err != nil {
			return err
		}

		if err := request.Validate(); err != nil {
			return err
		}

		if err := Broadcast(request, n); err != nil {
			return err
		}

		response := CreateOkResponse(request.Type, request.MsgID)
		return n.Reply(msg, response.ToMap())
	}
}
