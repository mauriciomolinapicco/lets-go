//unmarshal: convert JSON data into Go struct

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func main() {
	data := []byte(`
	{
		"id": 1,
		"name": "Mauricio",
		"email": "mauri@example.com",
		"active": true
	}
	`)

	var user User

	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}
