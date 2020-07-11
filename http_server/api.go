package httpServer

import (
	"github.com/tomiok/fuego-cache/http_server/operations"
	"net/http"
)

type Api struct {
	Server *FuegoHTTPServer
}

type Services struct {
	Ops *operations.OpsHandler
}

func NewHTTPApi(addr string, services Services) *Api {
	mux := http.NewServeMux()
	operations.AddRoutes(services.Ops, mux)

	return &Api{Server: NewHTTPServer(addr, mux)}
}

func (a *Api) Listen() {
	a.Server.Listen()
}
