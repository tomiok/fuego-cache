package cache

import "github.com/tomiok/fuego-cache/logs"

type g func(key interface{}) (string, error)
type a func(e Entry) string

//FuegoOps
type FuegoOps interface {
	Apply() FuegoResponse
}

type FuegoResponse struct {
	Response string
}

type WriteOperation struct {
	Operation string
	Key       string
	Value     string
	DoAdd     a
}

func (f *WriteOperation) Apply() FuegoResponse {
	e, err := ToEntry(f.Key, f.Value)

	if err != nil {
		logs.LogError(err)
		return FuegoResponse{Response: responseNil}
	}

	res := f.DoAdd(e)

	return FuegoResponse{
		Response: res,
	}
}

type ReadOperation struct {
	Operation string
	Key       string
	DoGet     g
}

func (r *ReadOperation) Apply() FuegoResponse {
	val, err :=	r.DoGet(r.Key)

	if err != nil {
		val = responseNil //todo fix this with an error response
	}

	return FuegoResponse{Response: val}
}
