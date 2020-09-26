package httpServer

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

func AddRoutes(o *OperationsHandler, mux *http.ServeMux) {
	mux.HandleFunc(GetUrl, o.GetValueHandler())
	mux.HandleFunc(SetUrl, o.SetValueHandler())
	mux.HandleFunc(DeleteUrl, o.DeleteValueHandler())
	// bulk operations
	mux.HandleFunc(BulkSetUrl, o.BulkSetHandler())
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
