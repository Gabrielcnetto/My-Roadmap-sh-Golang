package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()
	go perforTaks(ctx)
	time.Sleep(time.Second * 3)

}

func perforTaks(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Contexto cancelado")
		return
	}
}
