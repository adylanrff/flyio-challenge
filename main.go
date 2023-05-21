package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"flyio-challenge/app/broadcast/broadcast"
	broadcastRead "flyio-challenge/app/broadcast/read"
	broadcastStorage "flyio-challenge/app/broadcast/storage"
	broadcastTopology "flyio-challenge/app/broadcast/topology"

	"flyio-challenge/app/echo"
	"flyio-challenge/app/uuid"
)

func main() {
	initializeDependencies()

	n := maelstrom.NewNode()
	// Echo
	n.Handle("echo", echo.Handler(n))
	// UUID
	n.Handle("generate", uuid.Handler(n))

	// Broadcast
	n.Handle("broadcast", broadcast.Handler(n))
	n.Handle("read", broadcastRead.Handler(n))
	// ignore this one and use n.NodeIDs() instead
	n.Handle("topology", broadcastTopology.Handler(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func initializeDependencies() {
	broadcastStorage.InitMessageStorage()
}
