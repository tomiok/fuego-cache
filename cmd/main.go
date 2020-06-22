package main

import (
	"fmt"
	"github.com/tomiok/fuego-cache/tcp_server"
)

func main() {

	s := server.New("localhost:9919")
	s.OnNewMessage(func(c *server.Client, message string) {
		fmt.Println(message)
	})

	s.Listen()
}
