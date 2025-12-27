package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func numbers(ch chan<- int) {
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		ch <- i
	}

	defer wg.Done()

}

func alphabets(letters chan<- string) {
	for i := 'a'; i < 'g'; i++ {
		time.Sleep(100 * time.Millisecond)
		letters <- string(i)
	}

	defer wg.Done()
}

func main() {
	ints := make(chan int, 15)
	letters := make(chan string, 5)
	wg.Add(1)
	go numbers(ints)
	wg.Add(1)
	go alphabets(letters)
	for i := 0; i < 2; {
		select {
		case res1, ok := <-ints:
			if !ok {
				fmt.Println("Erro ao receber esse item do chanel")
			}
			fmt.Println(res1)
		case res2 := <-letters:
			fmt.Println(res2)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}

	defer wg.Wait()
	defer close(ints)
	defer close(letters)
}
