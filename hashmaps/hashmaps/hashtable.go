package main

import (
	"sync"
)

/*
struct que representa a hashtable
*/
type HashTable struct {
	/*
		cada lista de pessoas é guardada
		dentro de outra lista, que vai ser
		identificada por um int, que vai ser o
		hash, vamos criar um code hash pra salvar
		e dps pra buscar essa 'caixinha'
	*/
	Items map[int][]Pessoa
	lock  sync.RWMutex
}

func hash(nome string) (key int) {
	for _, letra := range nome {
		//cada letra, vira um iota(acho), ent cada letra tem
		// meio que 1-2-3, entao *31 monta uma hash completa
		key = 31*key + int(letra)
	}
	return
}

func (ht *HashTable) Put(pessoa Pessoa) int {
	//evitar concorrencia na hash table
	ht.lock.Lock()
	if len(ht.Items) == 0 {
		ht.Items = make(map[int][]Pessoa)
	}
	defer ht.lock.Unlock()
	key := hash(pessoa.Nome)
	ht.Items[key] = append(ht.Items[key], pessoa)
	return key
}

func (ht *HashTable) Remove(nome string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	key := hash(nome)
	delete(ht.Items, key)
}
func (ht *HashTable) Get(key int) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return ht.Items[key]
}

func (ht *HashTable) Search(nome string) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	key := hash(nome)
	return ht.Items[key]
}
