package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	type KeyValue struct {
		Key   string
		Value int
	}
	s := make([]KeyValue, 0, len(m))
	for k, v := range m {
		s = append(s, KeyValue{Key: k, Value: v})
	}

	//agora vamos fazer o sort
	/*
		 abaixo, a gente usa a funcao slicestable, passamos o s(lista que queremos usar)
		ai passamos uma funcao que recebe o i e o j, que sao os tipos de valores que queremos ordenar
		o i vai ser o primeiro elemento e o j o segundo elemento para compararmos item a item
	*/
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].Value < s[j].Value
	})
	for _, v := range s {
		fmt.Println(v.Key, "->", v.Value)
	}

}
