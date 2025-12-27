package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go performTask(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelando o for")
			return
		default:
			fmt.Println("Performando a task")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
