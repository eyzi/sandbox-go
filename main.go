package main

import (
	"encoding/json"
	"fmt"

	"github.com/eyzi/sandbox-go/server"
)

type OptionalObject struct {
	Extra string `json:"extra"`
}

type Object struct {
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Optional OptionalObject `json:"optional"`
}

func main() {
	jsonString := `{"name": "test-name", "age": 10, "extra": true}`

	var obj Object
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", obj)

	server.CreateServer()
}
