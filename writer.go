package duck

import (
	"net/http"
)

type Writer struct {
	http.ResponseWriter
	done chan bool
}

func newWriter(w http.ResponseWriter, req *http.Request) http.ResponseWriter {
	wrt := Writer{w, make(chan bool)}
	go func(wrt Writer) {
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

func (w Writer) Write(b []byte) (int, error) {
	w.done <- true
	return w.ResponseWriter.Write(b)
}

func (w Writer) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w Writer) WriteHeader(i int) {
	w.ResponseWriter.WriteHeader(i)
}
