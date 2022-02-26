package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gonotbook.com/24-mvc/models"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var u models.User

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = "001"

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.Write(uj)
}
