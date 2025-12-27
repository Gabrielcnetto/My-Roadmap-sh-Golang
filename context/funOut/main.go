package main

import (
	"fmt"
	"time"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func main() {
	fmt.Println("Start Fan Out ")
	c1 := generator(1, 2, 3)
	c2 := generator(4, 5, 6)

	go func() {
		for num := range c1 {
			fmt.Println(num)
		}
	}()
	go func() {
		for num := range c2 {
			fmt.Println(num)
		}
	}()

	time.Sleep(time.Second * 2)
}
