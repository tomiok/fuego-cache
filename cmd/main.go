package main

import (
	"github.com/tomiok/fuego-cache/http_server"
	"github.com/tomiok/fuego-cache/safe/fuego"
	"github.com/tomiok/fuego-cache/stdio_client"
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
		http := httpServer.NewHTTPServer("localhost:9919")
		http.Listen()
	} else {
		stdioClient.PrintBanner()
		stdioClient.StandardInputCache()
	}
}
