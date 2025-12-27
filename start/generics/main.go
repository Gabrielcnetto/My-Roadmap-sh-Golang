package main

import (
	"fmt"
)

type Number interface {
	~int32 | ~int64 | ~float32 | ~float64
}

/*
QUando a funcao tem operações matematicas, o generics é melhor usar com T
*/
func sumNumbers[T Number](numbers []T) T {
	var sum T
	for _, item := range numbers {
		sum += item
	}
	return sum
}

/*
Quando a funcao vai ter operações de leitura, ai é melhor usar um any
*/
type Container[T any] struct {
	items []T
}

func (c *Container[T]) Add(item T) {
	c.items = append(c.items, item)
}

func main() {
	ints32 := []int32{1, 2, 3, 4, 5}
	float64Val := []float64{2.2, 5.2, 100.2}
	fmt.Printf("\n ints32:%v\n", sumNumbers(ints32))
	fmt.Printf("\n float64Val:%v\n", sumNumbers(float64Val))
	//adicionando novos items
	listaAleatoria := map[string]interface{}{
		"Gabriel":    "Teste",
		"Wellington": "Teste2",
		"teste":      5,
	}
	container := Container[any]{}
	container.Add(listaAleatoria)
	fmt.Println(container.items...)
}
