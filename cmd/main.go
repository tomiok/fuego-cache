package main

import (
	"flag"
	"github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/http_server"
	"github.com/tomiok/fuego-cache/http_server/operations"
	"github.com/tomiok/fuego-cache/logs"
	"github.com/tomiok/fuego-cache/stdio_client"
	"github.com/tomiok/fuego-cache/tcp_server"
)

func main() {
	mode := flag.String("mode", "http", "Mode to run Fuego")
	flag.Parse()

	var fuegoInstance = cache.NewCache()
	if *mode == "tcp" {
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
	} else if *mode == "http" {
		addr := ":9919"
		api := httpServer.NewHTTPApi(addr, httpServer.Services{Ops: &operations.WebOperationsHandler{
			GetCallback: func(s interface{}) (string, error) {
				return fuegoInstance.GetOne(s)
			},
			SetCallback: func(k interface{}, v string, ttl int) (string, error) {
				return fuegoInstance.SetOne(k, v, ttl)
			},
			DeleteCallback: func(k interface{}) (string, error) {
				return fuegoInstance.DeleteOne(k), nil
			},
		}})
		logs.Info("stating server at " + addr)
		api.Listen()
	} else {
		s := stdioClient.NewStdClient()
		s.PrintBanner()
		s.OnNewMessage(func(message string) string {
			operationMessage := cache.NewFuegoMessage(message)
			ops := operationMessage.Compute(fuegoInstance)
			if ops != nil {
				return ops.Apply().Response
			}

			return "nil"
		})
		s.Listen()
	}
}
