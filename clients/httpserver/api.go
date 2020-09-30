package httpserver

import "net/http"

const (
	GetUrl    = "/fuego/get/"
	SetUrl    = "/fuego/set"
	DeleteUrl = "/fuego/del/"

	BulkSetUrl = "/fuego/bulk/set"

	// GET method is the default one
	DeleteMethod = "DELETE"
	PostMethod   = "POST"
)

// Web API structure

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

func AddRoutes(o *OperationsHandler, mux *http.ServeMux) {
	mux.HandleFunc(GetUrl, o.GetValueHandler())
	mux.HandleFunc(SetUrl, o.SetValueHandler())
	mux.HandleFunc(DeleteUrl, o.DeleteValueHandler())
	// bulk operations
	mux.HandleFunc(BulkSetUrl, o.BulkSetHandler())
}
