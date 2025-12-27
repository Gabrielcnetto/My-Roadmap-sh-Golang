package main

import "fmt"

func main() {
	var list = [4]int{1, 2, 3, 4}
	for _, item := range list[0:2] {
		fmt.Println(item)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i+1)
	}
	fmt.Println("Hello world")
}
