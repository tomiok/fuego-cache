package main

import (
	"fmt"
	"github.com/tomiok/fuego-cache/http_server"
	"github.com/tomiok/fuego-cache/logs"
	"github.com/tomiok/fuego-cache/tcp_server"
	"os"
)

func main() {
	mode := os.Getenv("MODE")

	if mode == "tcp" {
		s := server.New("localhost:9919")
		s.OnNewMessage(func(c *server.Client, message string) {
			fmt.Println(message)
		})

		s.Listen()
	} else if mode == "http" {
		http := httpserver.NewHTTPServer("localhost:9919")
		http.Listen()
	} else {
		logs.Fatal("Missing environment variable: MODE")
	}
}
