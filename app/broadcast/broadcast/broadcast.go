package broadcast

import (
	broadcastStorage "flyio-challenge/app/broadcast/storage"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

const MaxBroadcastRetry = 3

func Broadcast(request Request, n *maelstrom.Node) error {

	// Save message locally first
	messageStorage := broadcastStorage.GetMessageStore()
	messageStorage.SaveMessage(request.Message)

	// publish to other nodes with some retry...
	for _, nodeID := range n.NodeIDs() {
		if nodeID == n.ID() {
			continue
		}

		for i := 0; i < MaxBroadcastRetry; i++ {
			err := n.RPC(nodeID, request, func(msg maelstrom.Message) error {
				log.Println("ok")
				return nil
			})
			if err == nil {
				break
			}
		}
	}

	return nil
}
