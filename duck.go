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
	mutex.Lock()
	if data[req] == nil {
		data[req] = &context{values: make(map[interface{}]interface{})}
		mutex.Unlock()
	}
	mutex.Lock()
	data[req].values[key] = value
	mutex.Unlock()
	return
}

func GetContext(req *http.Request, key interface{}) interface{} {
	mutex.RLock()
	if data[req] == nil {
		return nil
	}
	value := data[req].values[key]
	mutex.RUnlock()
	return value
}

func GetAllContext(req *http.Request) map[interface{}]interface{} {
	mutex.RLock()
	if data[req] == nil {
		return nil
	}
	values := data[req].values
	mutex.RUnlock()
	return values
}

func DeleteContext(req *http.Request, key interface{}) {
	mutex.Lock()
	if data[req] != nil {
		delete(data[req].values, key)
	}
	mutex.Unlock()
}

func Clear(req *http.Request) {
	mutex.Lock()
	delete(data, req)
	mutex.Unlock()
}

func UseContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		defer Clear(req)
		next.ServeHTTP(rw, req)
	})
}
