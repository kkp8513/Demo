package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	
	student := &struct {
		Name string `json:"name,omitempty"`
		Age int `json:"Age"`
		Gender int `json:"-"`
		Phone int `json:",omitempty"`
	}{
		Name: "大哥",
		Age: 15,
		Gender: 1,
	}

	var res interface{}

	data, _ := json.Marshal(student)
	json.Unmarshal(data, &res)

	fmt.Println(string(data), res)
}

