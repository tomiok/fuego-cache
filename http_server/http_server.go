package httpServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type httpServer struct {
	address string
}

type HttpResponse struct {
	Value string `json:"value"`
}

type HttpRequestBody struct {
	Value string `json:"value"`
}

func getValue(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/fuego/")

	// TODO: get the actual stored value
	response := HttpResponse{}
	response.Value = id

	data, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func setValue(w http.ResponseWriter, r *http.Request) {
	var body HttpRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parsedBody, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Printf(string(parsedBody))

	w.Header().Set("Content-Type", "application/json")
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getValue(w, r)
	case "POST":
		setValue(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *httpServer) Listen() {
	http.HandleFunc("/fuego", httpHandler)

	http.ListenAndServe(s.address, nil)
}

func NewHTTPServer(address string) *httpServer {
	server := &httpServer{
		address: address,
	}

	return server
}
