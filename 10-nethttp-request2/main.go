package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type hotdog string

func (hd hotdog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	fmt.Fprintf(rw, "this is data", data)

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
