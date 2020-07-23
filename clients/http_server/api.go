package httpServer

import (
	"github.com/tomiok/fuego-cache/clients/http_server/web"
	"net/http"
)

type Api struct {
	Server *FuegoHTTPServer
}

type Services struct {
	Ops *web.OperationsHandler
}

func NewHTTPApi(addr string, services Services) *Api {
	mux := http.NewServeMux()
	web.AddRoutes(services.Ops, mux)

	return &Api{Server: NewHTTPServer(addr, mux)}
}

func (a *Api) Listen() {
	a.Server.Listen()
}
