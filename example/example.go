package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoo/duck"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)

	http.ListenAndServe(":8080", duck.UseContext(mux))
}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	rw = duck.NewWriter(rw, req)
	duck.SetContext(req, "key", req.RequestURI)
	fmt.Println("Context Size :", len(duck.GetAllContext(req)))
	rw.Write([]byte("Response writed " + duck.GetContext(req, "key").(string)))
	fmt.Println("Context Size :", len(duck.GetAllContext(req)))
}
