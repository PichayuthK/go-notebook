package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last  string
}

func main() {
	var p person
	s := `{"First": "peace","Last": "k"}`
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(p)
}
