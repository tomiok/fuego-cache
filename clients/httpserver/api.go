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

// custom handler generic type
type FuegoHandler func(w http.ResponseWriter, r *http.Request) error

func (f FuegoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		if err.Error() == "" {

		}
	}
}

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
	mux.Handle(GetUrl, FuegoHandler(o.GetValueHandler))
	mux.Handle(SetUrl, FuegoHandler(o.SetValueHandler))
	mux.Handle(DeleteUrl, FuegoHandler(o.DeleteValueHandler))
	// bulk operations
	mux.Handle(BulkSetUrl, FuegoHandler(o.BulkSetHandler))
}
