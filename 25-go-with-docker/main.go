package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	s := "this is index"
	rw.Write([]byte(s))
}
