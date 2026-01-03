package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var data = `
name: Jane Doe
age: 29
`

type Info struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	info := Info{}
	err := yaml.Unmarshal([]byte(data), &info)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", info)
}
