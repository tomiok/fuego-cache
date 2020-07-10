package httpServer

import (
	"github.com/tomiok/fuego-cache/http_server/operations"
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
)

type FuegoHTTPServer struct {
	address string
	s       http.Server
}



func (s *FuegoHTTPServer) Listen() {
	http.HandleFunc("/fuego", operations.httpHandler)

	logs.Fatal(http.ListenAndServe(s.address, nil))
}

func NewHTTPServer(address string) *FuegoHTTPServer {
	server := &FuegoHTTPServer{
		address: address,
	}

	return server
}
