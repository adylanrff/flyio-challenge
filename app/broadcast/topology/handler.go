package topology

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

		response := request
		response.Type = "topology_ok"
		return n.Reply(msg, response.ToMap())
	}
}
