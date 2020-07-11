package httpServer

import (
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
)

type FuegoHTTPServer struct {
	server *http.Server
}

func (f *FuegoHTTPServer) Listen() {
	logs.Fatal(f.server.ListenAndServe())
}

func NewHTTPServer(address string, mux http.Handler) *FuegoHTTPServer {
	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	return &FuegoHTTPServer{
		server: server,
	}
}
