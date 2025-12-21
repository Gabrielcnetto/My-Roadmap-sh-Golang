package main

import "fmt"

func main() {
	original := []int{1, 2, 3}
	copied := make([]int, 3, 10)
	copy(copied, original)
	fmt.Println(copied)
}
