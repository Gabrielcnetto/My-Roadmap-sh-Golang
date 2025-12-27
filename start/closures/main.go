package main

import "fmt"

func contador() func() int {
	value := 0
	return func() int {
		value += 1
		fmt.Println(value)
		return value
	}
}

func main() {
	cont := contador()
	cont()
	cont()
	cont()
	value := cont()
	fmt.Printf("value:%v", value)
}
