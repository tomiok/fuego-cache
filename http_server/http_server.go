package httpServer

import (
	"github.com/tomiok/fuego-cache/http_server/operations"
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
)

type FuegoHTTPServer struct {
	s *http.Server
	o *operations.OpsHandler
}

func (f *FuegoHTTPServer) Listen() {
	f.o.Routes()
	logs.Fatal(f.s.ListenAndServe())
}

func NewHTTPServer(address string, o *operations.OpsHandler) *FuegoHTTPServer {
	s := &http.Server{
		Addr: address,
	}

	return &FuegoHTTPServer{
		s: s,
		o: o,
	}
}
