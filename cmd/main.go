package main

import (
	"flag"
	"github.com/tomiok/fuego-cache/clients/http_server"
	"github.com/tomiok/fuego-cache/clients/http_server/web"
	"github.com/tomiok/fuego-cache/clients/stdio_client"
	"github.com/tomiok/fuego-cache/clients/tcp_server"
	"github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/logs"
)

func main() {
	mode := flag.String("mode", "http", "Mode to run Fuego")
	flag.Parse()

	var fuegoInstance = cache.NewCache()
	if *mode == "tcp" {
		s := tcpServer.New("localhost:9919")
		s.OnNewMessage(func(c *tcpServer.Client, message string) {
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
		api := httpServer.NewHTTPApi(addr, httpServer.Services{Ops: &web.OperationsHandler{
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
