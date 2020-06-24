package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/tomiok/fuego-cache/logs"
	"net/http"
	"strings"
)

type httpServer struct {
	address string
}

type HttpResponse struct {
	Data string `json:"data"`
}

type HttpRequestBody struct {
	Value string `json:"value"`
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/fuego/")

	response := HttpResponse{}
	response.Data = id

	data, err := json.Marshal(response)

	if err != nil {
		logs.Error("Error parsing json: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received apost")

	var body HttpRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		logs.Error("Error decoding body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	parsedBody, err := json.Marshal(body)

	if err != nil {
		logs.Error("Error parsing json: " + err.Error())
	}

	fmt.Printf(string(parsedBody))

	w.Header().Set("Content-Type", "application/json")
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	}
}

func (s *httpServer) Listen() {
	http.HandleFunc("/fuego/", httpHandler)

	http.ListenAndServe(s.address, nil)
}

func NewHTTPServer(address string) *httpServer {
	server := &httpServer{
		address: address,
	}

	return server
}
