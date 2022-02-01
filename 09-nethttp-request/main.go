package main

import (
	"fmt"
	"log"
	"net/http"
)

type myRequest string

func (mr myRequest) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(rw, "This is body %v", req.PostForm)
}

func main() {
	var mr myRequest
	http.ListenAndServe(":8080", mr)
}
