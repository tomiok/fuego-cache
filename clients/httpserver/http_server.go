package httpserver

import (
	"context"
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type FuegoHTTPServer struct {
	server *http.Server
}

func (f *FuegoHTTPServer) Listen() {
	go func() {
		if err := f.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatal("could not listen " + f.server.Addr + " due to: " + err.Error())
		}
	}()
	f.gracefulShutdown()
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

func (f *FuegoHTTPServer) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	logs.Info("server is shutting down " + sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	f.server.SetKeepAlivesEnabled(false)
	if err := f.server.Shutdown(ctx); err != nil {
		logs.Fatal("could not gracefully shutdown the server " + err.Error())
	}
	logs.Info("server stopped")
}
