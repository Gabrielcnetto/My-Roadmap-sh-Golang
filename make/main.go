package main

import "fmt"

func main() {
	//make serve para 3 tipos, channels, slice e maps
	nums := make([]int, 0, 100) //caso eu saber quantos items eu vou usar, nao usa 3 atributos usa apenas 2: make([]int, 100)
	//caso eu nao saber quantos itens eu vou ter, ai começamos com 0: make([]int, 0, 100)
	fmt.Println(nums)
}
