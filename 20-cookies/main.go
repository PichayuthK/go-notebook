package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", setCookie)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}

func setCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value cookie",
	})
	fmt.Fprintln(w, "Cookie written")
	fmt.Fprintln(w, "check in dev tools")
}

func readCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprintln(w, "This is cookie:", c)
}
