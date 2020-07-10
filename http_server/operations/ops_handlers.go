package operations

import (
	"encoding/json"
	"github.com/tomiok/fuego-cache/internal"
	"net/http"
	"strings"
)

type OpsHandler struct {
	GetCallback func(s string) string
	SetCallback func(k, v string) string
}

func (o *OpsHandler) GetValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/fuego/")

	res := o.GetCallback(id)

	_ = json.NewEncoder(w).Encode(GetResponse{Response: res})
}

func (o *OpsHandler) SetValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	var b SetRequest
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res := o.SetCallback(b.Key, b.Value)

	internal.OnCloseError(body.Close)
	_ = json.NewEncoder(w).Encode(&res)
}

func (o *OpsHandler) DeleteValueHandler(w http.ResponseWriter, r *http.Request) {

}

func (o *OpsHandler) Routes() {
	http.HandleFunc("/fuego/get/", o.GetValueHandler)
	http.HandleFunc("/fuego/set/", o.SetValueHandler)
	http.HandleFunc("/fuego/del/", o.DeleteValueHandler)
}

type GetResponse struct {
	Response string `json:"response"`
}

type SetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
