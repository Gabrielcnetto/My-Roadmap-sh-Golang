package main

import (
	"fmt"
	"sync"
)

/*
Principais componentes de uma worker pool
1) Contém uma lista de processos que precisam ser feitos, tipo uma fila
2) *Worker goroutines*: Um número fixo de goroutines
que continuamente escutam novos trabalhos na fila de trabalhos e os processam.
3) *Coletor de resultado*: Geralmente outra goroutine que processa os resultados dos trabalhadores
4) *Despachante*: Componente que organiza os trabalhos e envia cada item da fila
para determinado goroutine executar, incluindo sincronização e desligamento
5) *Mecanismos de sincronização*: sync.waitGroup
*/
type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Square int
}

// aqui embaixo tem o cara que trabalha
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{JobID: job.ID, Square: job.Value * job.Value}
	}
}

// aqui o cara que coleta os resultados, e organiza
func collectResults(results <-chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		fmt.Printf("Job ID: %d, Input: %d, Squared Value: %d\n", result.JobID, result.JobID, result.Square)
	}
}

// aqui o gerente
func dispatcher(jobCount, workerCount int) {
	jobs := make(chan Job, jobCount)
	results := make(chan Result, jobCount)

	var wg sync.WaitGroup

	// Start workers
	wg.Add(workerCount)
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results, &wg)
	}

	// Start collecting results
	var resultsWg sync.WaitGroup
	resultsWg.Add(1)
	go collectResults(results, &resultsWg)

	// Distribute jobs and wait for completion
	for j := 1; j <= jobCount; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs)
	wg.Wait()
	close(results)

	// Ensure all results are collected
	resultsWg.Wait()
}
func main() {
	const jobCount = 100  // Total number of jobs to process
	const workerCount = 3 // Number of workers to process the jobs

	fmt.Println("Starting batch processing with synchronized result collection...")
	dispatcher(jobCount, workerCount)
}
