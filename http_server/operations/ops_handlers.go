package operations

import (
	"encoding/json"
	"fmt"
	"github.com/tomiok/fuego-cache/internal"
	"net/http"
	"strings"
)

type OpsHandler struct {
	GetCallback func(s string) string
	SetCallback func(k, v string) string
}

func (o *OpsHandler) GetValueHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/fuego/")

	res := o.GetCallback(id)

	data, err := json.Marshal(struct {
		Response string `json:"response"`
	}{Response: res})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}

func (o *OpsHandler) SetValueHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parsedBody, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Printf(string(parsedBody))
	internal.OnCloseError(body.Close)
	w.Header().Set("Content-Type", "application/json")
}

func DeleteValueHandler(w http.ResponseWriter, r *http.Request) {

}
