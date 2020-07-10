package httpServer

import "github.com/tomiok/fuego-cache/http_server/operations"

type Api struct {
	Server *FuegoHTTPServer
}

type Services struct {
	Ops *operations.OpsHandler
}

func NewHTTPApi(addr string, s Services) *Api {
	mux :=	operations.Routes(s.Ops)
	return &Api{Server: NewHTTPServer(addr, mux)}
}

func (a *Api) Start() {
	a.Server.Listen()
}
