package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	var p Person

	jdata := `{"name": "Mr Smith", "age": 45}`

	dec := json.NewDecoder(strings.NewReader(jdata))

	dec.Decode(&p)

	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
