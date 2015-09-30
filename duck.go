package duck

import (
	"net/http"
	"sync"
)

type context struct {
	values map[interface{}]interface{}
}

var (
	mutex sync.RWMutex
	data  = make(map[*http.Request]*context)
)

func SetContext(req *http.Request, key, value interface{}) {
	if data[req] == nil {
		data[req] = &context{values: make(map[interface{}]interface{})}
	}
	mutex.Lock()
	data[req].values[key] = value
	mutex.Unlock()
	return
}

func GetContext(req *http.Request, key interface{}) interface{} {
	if data[req] == nil {
		return nil
	}
	return data[req].values[key]
}

func GetAllContext(req *http.Request) map[interface{}]interface{} {
	if data[req] == nil {
		return nil
	}
	return data[req].values
}

func DeleteContext(req *http.Request, key interface{}) {
	mutex.Lock()
	if data[req] != nil {
		delete(data[req].values, key)
	}
	mutex.Unlock()
}

func Clear(req *http.Request) {
	delete(data, req)
}

func UseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer Clear(req)
		next.ServeHTTP(rw, req)
	})
}
