package operations

import (
	"encoding/json"
	"github.com/tomiok/fuego-cache/internal"
	"net/http"
	"strings"
)

type WebOperationsHandler struct {
	GetCallback    func(k interface{}) (string, error)
	SetCallback    func(k interface{}, v string) (string, error)
	DeleteCallback func(k interface{}) (string, error)
}

func (o *WebOperationsHandler) GetValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := strings.TrimPrefix(r.URL.Path, GetUrl)
		var _id interface{}
		_id = id
		res, err := o.GetCallback(_id)

		_ = json.NewEncoder(w).Encode(WebResponse{Response: res, Err: err != nil})
	}
}

func (o *WebOperationsHandler) SetValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == PostMethod {
			body := r.Body
			var b SetRequest
			err := json.NewDecoder(body).Decode(&b)

			if err != nil || b.Key == "" {
				http.Error(w, "cannot process current request", http.StatusBadRequest)
				return
			}

			res, err := o.SetCallback(b.Key, b.Value)

			internal.OnCloseError(body.Close)
			_ = json.NewEncoder(w).Encode(WebResponse{Response: res, Err: err != nil})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func (o *WebOperationsHandler) DeleteValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == DeleteMethod {
			id := strings.TrimPrefix(r.URL.Path, DeleteUrl)
			deleted, err := o.DeleteCallback(id)
			_ = json.NewEncoder(w).Encode(WebResponse{Response: deleted, Err: err != nil})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func AddRoutes(o *WebOperationsHandler, mux *http.ServeMux) {
	mux.HandleFunc(GetUrl, o.GetValueHandler())
	mux.HandleFunc(SetUrl, o.SetValueHandler())
	mux.HandleFunc(DeleteUrl, o.DeleteValueHandler())
}

type WebResponse struct {
	Response string `json:"response"`
	Err      bool   `json:"err"`
}

type SetRequest struct {
	Key   interface{} `json:"key"`
	Value string      `json:"value"`
}
