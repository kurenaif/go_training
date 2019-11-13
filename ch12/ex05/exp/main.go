package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
	C    int
	B    bool
	S    []string
	P    *string
}

func main() {
	human := Human{
		Name: "Taro",
		Age:  0,
		C:    1,
		B:    true,
		S:    []string{"Hello", "World"},
		P:    nil,
	}

	bytes, _ := json.Marshal(human)
	string := string(bytes)
	fmt.Println(string)
}
