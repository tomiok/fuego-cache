package httpserver

import (
	"encoding/json"
	"errors"
	cache "github.com/tomiok/fuego-cache/fuego"
	"github.com/tomiok/fuego-cache/internal"
	"net/http"
	"strings"
)

type OperationsHandler struct {
	GetCallback    func(k string) (string, error)
	SetCallback    func(k string, v string, ttl int) (string, error)
	DeleteCallback func(k string) (string, error)
	//bulks operations
	BulkGetCallback    func(keys []string)
	BulkSetCallback    func(bulkEntry cache.BulkEntry) cache.BulkResponse
	BulkDeleteCallback func(keys []string)
}

func (o *OperationsHandler) GetValueHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	key := strings.TrimPrefix(r.URL.Path, GetUrl)
	res, err := o.GetCallback(key)
	//when a response is with error true and value is nil, it means that the key is not present in the cache
	_ = json.NewEncoder(w).Encode(HTTPResponse{Response: res, Err: err != nil})
	return nil
}

func (o *OperationsHandler) SetValueHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == PostMethod {
		body := r.Body
		var b SetEntryCMD
		err := json.NewDecoder(body).Decode(&b)

		defer func() {
			w.WriteHeader(500)

			if r := recover(); r != nil {
				_ = json.NewEncoder(w).Encode(HTTPError{Msg: "ERROR DUDE"})
			}

		}()

		if err != nil || b.Key == "" {
			return err
		}

		res, err := o.SetCallback(b.Key, b.Value, b.TTL)
		internal.OnCloseError(body.Close)
		_ = json.NewEncoder(w).Encode(HTTPResponse{Response: res, Err: err != nil})
		return nil
	} else {
		return errors.New("methodNotAllowed")
	}
}

func (o *OperationsHandler) DeleteValueHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == DeleteMethod {
		id := strings.TrimPrefix(r.URL.Path, DeleteUrl)
		deleted, err := o.DeleteCallback(id)
		_ = json.NewEncoder(w).Encode(HTTPResponse{Response: deleted, Err: err != nil})
		return nil
	} else {
		return errors.New("methodNotAllowed")
	}
}

func (o *OperationsHandler) BulkSetHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == PostMethod {
		var cmd []BulkSetCMD
		body := r.Body
		defer internal.OnCloseError(body.Close)
		_ = json.NewDecoder(body).Decode(&cmd)
		res := o.BulkSetCallback(toBulkEntry(cmd))
		_ = json.NewEncoder(w).Encode(&res)

		return nil
	} else {
		return errors.New("methodNotAllowed")
	}
}

// HTTP response and request helpers
type SetEntryCMD struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int    `json:"ttl,omitempty"` //if 0 it is supposed no TTL, those are IN SECONDS
}

type BulkSetCMD struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int    `json:"ttl,omitempty"`
}

type HTTPResponse struct {
	Response string `json:"response"`
	Err      bool   `json:"err"`
}

type HTTPError struct {
	Msg string `json:"message"`
}

func toBulkEntry(body []BulkSetCMD) cache.BulkEntry {
	var bulkEntry cache.BulkEntry
	for _, entry := range body {
		bulkEntry.Add(entry.Key, entry.Value, entry.TTL)
	}
	return bulkEntry
}
