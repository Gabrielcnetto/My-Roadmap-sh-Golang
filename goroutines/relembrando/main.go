package main

import (
	"fmt"
	"sync"
)

func main() {
	list := []float64{50, 30, 15, 25}
	chant := make(chan float64)
	var wg sync.WaitGroup

	wg.Add(len(list))
	for index, item := range list {
		go func(i int, v float64) {
			defer wg.Done()
			chant <- float64(i) * v
		}(index, item)
	}

	//wg wait a gente espera em outro goroutine, feita apenas pra esperar
	go func() {
		wg.Wait()
		close(chant)
	}()
	var sum float64
	for v := range chant {
		sum += v
	}
	fmt.Printf("soma: %v\n", sum)
}
