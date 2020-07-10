package httpServer

type Api struct {
	Server FuegoHTTPServer
}

func (a *Api) Start() {
	a.Server.Listen()
}