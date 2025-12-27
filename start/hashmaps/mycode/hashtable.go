package main

import "sync"

type Hashcode struct {
	Lista map[int][]Carro
	Lock  sync.RWMutex
}

func getHash(marca string) (key int) {
	for _, letra := range marca {
		key = key*31 + int(letra)
	}
	return
}

func (ht *Hashcode) Put(carro Carro) (key int) {

	hashDaMarca := getHash(carro.Marca)
	if ht.Lista == nil {
		ht.Lista = make(map[int][]Carro)
	}
	ht.Lista[hashDaMarca] = append(ht.Lista[hashDaMarca], carro)
	return hashDaMarca
}

func (ht *Hashcode) GetByBrand(brand string) []Carro {
	ht.Lock.Lock()
	defer ht.Lock.RUnlock()
	brandHash := getHash(brand)
	return ht.Lista[brandHash]
}
