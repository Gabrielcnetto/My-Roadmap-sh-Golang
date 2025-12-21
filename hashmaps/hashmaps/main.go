package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Sexo      string
	Idade     int
}

func main() {
	pessoas := []Pessoa{
		Pessoa{Nome: "Maria", Sobrenome: "Graça", Sexo: "F", Idade: 55},
		Pessoa{Nome: "Gabriel", Sobrenome: "Netto", Sexo: "M", Idade: 21},
		Pessoa{Nome: "Wellington", Sobrenome: "Netto", Sexo: "M", Idade: 28},
		Pessoa{Nome: "Wellington C.", Sobrenome: "Netto", Sexo: "M", Idade: 28},
	}
	table := HashTable{}
	keys := make([]int, len(pessoas))
	for i, pessoas := range pessoas {
		keys[i] = table.Put(pessoas)
	}
	for _, pessoa := range keys {
		ps := table.Get(pessoa)
		for _, item := range ps {
			fmt.Println(item.Nome, item.Sobrenome)
		}
	}
	finded := table.Search("Wellington")
	fmt.Println(finded)
}
