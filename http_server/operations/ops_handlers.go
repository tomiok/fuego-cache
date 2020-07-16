package operations

import (
	"encoding/json"
	"fmt"
	"github.com/tomiok/fuego-cache/internal"
	"net/http"
	"strings"
)

const baseURL = "/fuego"

type OpsHandler struct {
	GetCallback func(s string) string
	SetCallback func(k, v string) string
}

func (o *OpsHandler) GetValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := strings.TrimPrefix(r.URL.Path, GetUrl)

		res := o.GetCallback(id)

		_ = json.NewEncoder(w).Encode(GetResponse{Response: res})
	}
}

func (o *OpsHandler) SetValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "POST" {
			body := r.Body
			var b SetRequest
			err := json.NewDecoder(body).Decode(&b)

			if err != nil || b.Key == "" {
				http.Error(w, "cannot process current request", http.StatusBadRequest)
				return
			}

			res := o.SetCallback(b.Key, b.Value)

			internal.OnCloseError(body.Close)
			_ = json.NewEncoder(w).Encode(&res)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return

		}
	}
}

func (o *OpsHandler) DeleteValueHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func AddRoutes(o *OpsHandler, mux *http.ServeMux) {
	mux.HandleFunc(GetUrl, o.GetValueHandler())
	mux.HandleFunc(SetUrl, o.SetValueHandler())
	mux.HandleFunc(DeleteUrl, o.DeleteValueHandler())
}

type GetResponse struct {
	Response string `json:"response"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func urlGen(urlPath string) string {
	return fmt.Sprintf("%s/%s", baseURL, urlPath)
}
