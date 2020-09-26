package main

import (
	"fmt"
	"github.com/tomiok/fuego-cache/clients/httpserver"
	"github.com/tomiok/fuego-cache/clients/stdioclient"
	"github.com/tomiok/fuego-cache/clients/tcpserver"
	"github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/logs"
)

func main() {
	config := cache.ParseConfiguration()

	var fuegoInstance = cache.NewCache(config)
	if config.Mode == "tcp" {
		s := tcpServer.New("localhost:9919")
		s.OnNewMessage(func(c *tcpServer.Client, message string) {
			operationMessage := cache.NewFuegoMessage(message)
			ops, err := operationMessage.Compute(fuegoInstance)
			if err != nil {
				return
			}
			response := ops.Apply()
			_ = c.Send(response.Response + "\n")
		})

		s.Listen()
	} else if config.Mode == "http" {
		addr := fmt.Sprintf(":%s", config.WebPort)
		api := httpServer.NewHTTPApi(addr, httpServer.Services{Ops: &httpServer.OperationsHandler{
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
			ops, err := operationMessage.Compute(fuegoInstance)
			if err != nil {
				return operationMessage.ErrResponse
			}
			return ops.Apply().Response

		})
		s.Listen()
	}
}
