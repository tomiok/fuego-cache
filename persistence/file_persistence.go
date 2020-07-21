package persistence

import "time"

type Persist interface {
	Save(operation string, k interface{}, value string, ts int64)
}

func Apply(p Persist, data Data) {
	p.Save(data.operation, data.key, data.value, time.Now().Unix())
}

type Data struct {
	operation string
	key       interface{}
	value     string
}

type FilePersistence struct {
	File string
}

func (f *FilePersistence) Save(operation string, k interface{}, value string, ts int64) {

}
