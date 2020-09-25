package httpServer

import (
	"net/http"
)

type Api struct {
	Server *FuegoHTTPServer
}

type Services struct {
	Ops *OperationsHandler
}

func NewHTTPApi(addr string, services Services) *Api {
	mux := http.NewServeMux()
	AddRoutes(services.Ops, mux)

	return &Api{Server: NewHTTPServer(addr, mux)}
}

func (a *Api) Listen() {
	a.Server.Listen()
}
