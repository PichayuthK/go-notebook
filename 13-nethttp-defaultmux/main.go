package main

import (
	"fmt"
	"net/http"
)

func d(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "this is my response!")
}

func main() {
	http.HandleFunc("/hot", d)

	http.ListenAndServe(":8080", nil)
}
