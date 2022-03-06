package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", getBooks).Schemes("http").Methods("GET")

	author := r.PathPrefix("/authors").Subrouter()
	author.HandleFunc("", getAuthors)

	http.ListenAndServe(":8080", r)
}

func getBooks(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Println(vars["title"])
	fmt.Println(vars["page"])
	json.NewEncoder(rw).Encode(vars)
}

func getAuthors(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "getting all authors")
}
