package main

import "fmt"

func result(data interface{}) float64 {
	value, ok := data.(float64)
	if !ok {
		value = 0
	}
	fmt.Println(value)
	return value * 45
}

func main() {
	var i interface{}
	i = 42.4
	fmt.Println(result(i))
}

func describe(i interface{}) {
	fmt.Printf("Valor:%v", i)
}
