package main

import "fmt"

func main() {
	//aqui depende do valor do iota, tipo o 1 faz *2
	// entao o A vai ser index 0 *2 =0
	// index B vai ser 1 *2 = 2
	// index c vai ser 2 *2 = 4
	const (
		A = iota << 1
		B
		C
		D
		E
	)
	fmt.Println(D) //printou 6
	const (
		u = iota * 42 // u == 0     (untyped integer constant)
		Z
	)
	fmt.Println(Z)
}
