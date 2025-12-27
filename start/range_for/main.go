package main

import (
	"fmt"
	"sync"
)

func main() {
	//	var mutex = &sync.RWMutex{}
	var wg = &sync.WaitGroup{}
	msg1 := make(chan string, 5)
	msg2 := make(chan string, 5)
	wg.Add(1)

	go func() {
		msg1 <- "Olá, aqui é mensagem 1"
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		msg2 <- "Olá, aqui é mensagem 2"
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-msg1:
			fmt.Println("Received", msg1)
		case msg2 := <-msg2:
			fmt.Println("Received", msg2)
		}
	}
	wg.Wait()

}
