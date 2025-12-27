package main

import (
	"fmt"
	"sync"
)

type Empregado struct {
	Id      int
	Salario float64
}
type Result struct {
	Empregado
}

func worker(id int, trabalhadores <-chan Empregado, resultado chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range trabalhadores {
		resultado <- Result{Empregado: Empregado{Id: job.Id, Salario: job.Salario * 1.12}}
	}
}

func Leitor(resultados <-chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range resultados {
		fmt.Printf("\nID:%v - Salario:%v\n", result.Id, result.Salario)
	}
}

func gerente(TotalDeTrabalhadores, workers int) {
	trabalhadores := make(chan Empregado, TotalDeTrabalhadores)
	resultado := make(chan Result, TotalDeTrabalhadores)
	var wg sync.WaitGroup
	//iniciamos 3 goroutines pra irem executando, adiciona em for (cria 3 goroutines de worker)
	wg.Add(workers)
	for w := 1; w <= workers; w++ {
		go worker(w, trabalhadores, resultado, &wg)
	}
	var resultWg sync.WaitGroup // vamos iniciar outro wg apenas para cuidar do resultado, pra nao bagunçar
	resultWg.Add(1)
	go Leitor(resultado, &resultWg) //aqui ele fica escutando com um for tudo que cai em resultado, e fecha sempre que termina

	//pra tudo isso funcionar, precisamos injetar, e faremos isso com for:
	for j := 0; j < TotalDeTrabalhadores; j++ {
		trabalhadores <- Empregado{Id: j, Salario: 5 + float64(j)}
	}
	close(trabalhadores)
	wg.Wait()
	close(resultado)
	resultWg.Wait()

}
func main() {
	const (
		TotalDeTrabalhadores = 100
		workers              = 3
	)
	gerente(TotalDeTrabalhadores, workers)
}
