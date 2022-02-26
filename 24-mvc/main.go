package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gonotbook.com/24-mvc/controllers"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController()

	r.GET("/", index)
	r.POST("/user", uc.CreateUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := "this is index"
	w.Write([]byte(s))
}
