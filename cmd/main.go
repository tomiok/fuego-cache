package main

import (
	cache "github.com/tomiok/fuego-cache/safe/fuego"
	"github.com/tomiok/fuego-cache/tcp_server"
)

func main() {

	var fuegoInstance = cache.NewCache()

	s := server.New("localhost:9919")
	s.OnNewMessage(func(c *server.Client, message string) {
		operationMessage := cache.NewFuegoMessage(message)
		ops := operationMessage.Compute(fuegoInstance)
		ops.Apply()
	})

	s.Listen()
}
