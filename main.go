package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}

	// Marshal returns a byte slice and an error
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response

	// unmarshal takes a json encoded byte slice, pointer where to put it
	// decodes the json
	_ = json.Unmarshal(j, &r2)
	fmt.Printf("%#v\n", r2)
}
