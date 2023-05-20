package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"

	"flyio-challenge/app/echo"
	"flyio-challenge/app/uuid"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("echo", echo.Handler(n))
	n.Handle("generate", uuid.Handler(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
