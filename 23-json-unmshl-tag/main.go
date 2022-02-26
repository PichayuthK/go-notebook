package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"firstname"`
	Last  string `json:"lastname"`
}

func main() {
	var px []person
	s := `[{"firstname":"Peace", "lastname": "K"}, {"firstname":"John", "lastname":"Doe"}]`

	err := json.Unmarshal([]byte(s), &px)
	if err != nil {
		log.Println(err)
	}

	for _, p := range px {
		fmt.Println(p)
	}
}
