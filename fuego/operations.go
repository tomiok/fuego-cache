package cache

type getCallback func(key interface{}) (string, error)
type addCallback func(key interface{}, value string, ttl ...int) (string, error)

//FuegoOps
type FuegoOps interface {
	Apply() FuegoResponse
}

type FuegoResponse struct {
	Response string
	Err      bool
}

type WriteOperation struct {
	Operation string
	Key       string
	Value     string
	DoAdd     addCallback
}

func (f *WriteOperation) Apply() FuegoResponse {
	res, err := f.DoAdd(f.Key, f.Value)

	return FuegoResponse{
		Response: res,
		Err:      err != nil,
	}
}

type ReadOperation struct {
	Operation string
	Key       string
	DoGet     getCallback
}

func (r *ReadOperation) Apply() FuegoResponse {
	val, err := r.DoGet(r.Key)

	if err != nil {
		val = responseNil //todo fix this with an error response
	}

	return FuegoResponse{Response: val}
}
