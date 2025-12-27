package main

import (
	"fmt"
	"sync"
	"time"
)

/*
É usado quando uma única função lê várias entradas e
prossegue até que todas sejam fechadas.
Isto é possível multiplexando a entrada em um único canal.
*/

func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(in))
	for _, c := range in {
		go output(c)
	}
	go func() {
		wg.Wait()
	}()
	return out
}
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
	c1 := generator(1, 2, 3)
	c2 := generator(4, 5, 6)
	c3 := merge(c1, c2)
	go func() {
		for num := range c3 {
			fmt.Println("____________")
			fmt.Println(num)
		}
	}()
	time.Sleep(time.Second * 5)
}
