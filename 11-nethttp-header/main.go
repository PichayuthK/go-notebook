package main

import (
	"fmt"
	"net/http"
)

type hotdog string

func (hd hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Peace", "Peace info")
	fmt.Fprintln(w, "this is in body")
}

func main() {
	var hd hotdog
	http.ListenAndServe(":8080", hd)
}
