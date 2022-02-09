package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (hd hotdog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "this is my response")
}

func main() {
	var d hotdog

	mux := http.NewServeMux()
	mux.Handle("/hot", d)

	http.ListenAndServe(":8080", mux)
}
