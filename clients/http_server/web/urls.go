package web

import "net/http"

const (
	GetUrl    = "/fuego/get/"
	SetUrl    = "/fuego/set"
	DeleteUrl = "/fuego/del/"

	DeleteMethod = "DELETE"
	PostMethod   = "POST"
)

func AddRoutes(o *OperationsHandler, mux *http.ServeMux) {
	mux.HandleFunc(GetUrl, o.GetValueHandler())
	mux.HandleFunc(SetUrl, o.SetValueHandler())
	mux.HandleFunc(DeleteUrl, o.DeleteValueHandler())
}

type HTTPRes struct {
	Response string `json:"response"`
	Err      bool   `json:"err"`
}

type SetRequest struct {
	Key   interface{} `json:"key"`
	Value string      `json:"value"`
	TTL   int         `json:"ttl,omitempty"` //if 0 it is supposed no TTL, those are IN SECONDS
}
