package duck

import (
	"net/http"
)

type writer struct {
	http.ResponseWriter
	done chan bool
}

func NewWriter(rw http.ResponseWriter, req *http.Request) http.ResponseWriter {
	wrt := writer{rw, make(chan bool)}
	go func(wrt writer) {
		for {
			select {
			case <-wrt.done:
				Clear(req)
				return
			default:
				continue
			}
		}
	}(wrt)
	return wrt
}

func (w writer) Write(b []byte) (int, error) {
	w.done <- true
	return w.ResponseWriter.Write(b)
}

func (w writer) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w writer) WriteHeader(i int) {
	w.ResponseWriter.WriteHeader(i)
}
