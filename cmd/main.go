package main

import (
	"bufio"
	httpserver "github.com/tomiok/fuego-cache/http_server"
	"github.com/tomiok/fuego-cache/internal"
	"github.com/tomiok/fuego-cache/logs"
	cache "github.com/tomiok/fuego-cache/safe/fuego"
	"os/signal"

	"github.com/tomiok/fuego-cache/tcp_server"
	"os"
)

func main() {
	printBanner()
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
		standardInputCache()
	}
}

func standardInputCache() {
	quit := make(chan os.Signal, 1)
	go func() {
		for {
			logs.Info("start with fuego here... (add 1 1)")
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			logs.Info(text)
		}
	}()

	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		logs.Info("exiting...")
	}
}

func printBanner() {
	internal.PrintBanner()
}
