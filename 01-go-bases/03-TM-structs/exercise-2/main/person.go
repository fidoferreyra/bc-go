package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type Employee struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Person   `json:"personal_data"`
}

func (e Employee) Print() {
	eJSON, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(eJSON))
	}
}
