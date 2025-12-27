package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello 1")
}
func main() {
	go hello() //goroutine roda imediatamente, mas o main nao espera nada dele, entao forçamos o time sleep pra dar tempo pra ele falar algo
	time.Sleep(1 * time.Second)
	fmt.Println("Hello 2")
}
