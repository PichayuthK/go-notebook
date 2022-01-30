package main

import (
	"fmt"
	"net/http"
)

type myServer int

func (m myServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "my test response")
}

func main() {
	var s myServer

	http.ListenAndServe(":8080", s)
}
