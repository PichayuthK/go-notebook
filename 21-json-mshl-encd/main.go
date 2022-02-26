package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
	Last  string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	s := "this is foo."
	rw.Write([]byte(s))
}

func mshl(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	rw.Write(json)
}

func encd(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	s := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun"},
	}
	err := json.NewEncoder(rw).Encode(s)

	if err == nil {
		log.Println(err)
	}

}
