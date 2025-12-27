package main

import "fmt"

type Carro struct {
	Marca    string
	Potencia float64
	IPVA     float64
}

func main() {
	carros := []Carro{
		Carro{Marca: "VW", Potencia: 1.0, IPVA: 894.90},
		Carro{Marca: "VW", Potencia: 1.5, IPVA: 992.30},
		Carro{Marca: "Ford", Potencia: 1.7, IPVA: 773.29},
	}
	table := Hashcode{}
	Salvos := map[int]Carro{}
	for index, item := range carros {
		table.Put(item)
		Salvos[index] = item
	}
	fmt.Println(table.Lista)
}
