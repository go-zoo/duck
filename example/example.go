package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/duck"
)

func main() {
	http.Handle("/", duck.Watch(testHandler))

	http.ListenAndServe(":8080", nil)
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	duck.SetContext(req, "key", "value")
	fmt.Println(duck.GetContext(req, "key").(string))
	rw.Write([]byte("Response writed"))
	fmt.Println(len(duck.GetAllContext(req)))
}
