package httpServer

import "github.com/tomiok/fuego-cache/http_server/operations"

type Api struct {
	Server *FuegoHTTPServer
}

func NewHTTPApi(addr string, o *operations.OpsHandler) *Api{
	return &Api{Server:  NewHTTPServer(addr, o)}
}

func (a *Api) Start() {
	a.Server.Listen()
}
