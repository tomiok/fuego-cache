package main

import (
	httpserver "github.com/tomiok/fuego-cache/http_server"
	"github.com/tomiok/fuego-cache/logs"
	cache "github.com/tomiok/fuego-cache/safe/fuego"

	"github.com/tomiok/fuego-cache/tcp_server"
	"os"
)

func main() {
	mode := os.Getenv("MODE")
	var fuegoInstance = cache.NewCache()
	if mode == "tcp" {
		s := server.New("localhost:9919")
		s.OnNewMessage(func(c *server.Client, message string) {
			operationMessage := cache.NewFuegoMessage(message)
			ops := operationMessage.Compute(fuegoInstance)
			if ops != nil {
				response := ops.Apply()
				_ = c.Send(response.Response + "\n")
			}
		})

		s.Listen()
	} else if mode == "http" {
		http := httpserver.NewHTTPServer("localhost:9919")
		http.Listen()
	} else {
		logs.Fatal("Missing environment variable: MODE")
	}
}
