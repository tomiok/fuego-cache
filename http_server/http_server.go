package httpServer

import (
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
)

type FuegoHTTPServer struct {
	s *http.Server
}

func (f *FuegoHTTPServer) Listen() {
	logs.Fatal(f.s.ListenAndServe())
}

func NewHTTPServer(address string, mux http.Handler) *FuegoHTTPServer {
	s := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return &FuegoHTTPServer{
		s: s,
	}
}
