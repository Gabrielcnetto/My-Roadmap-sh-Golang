package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	userIDs := []int{101, 102, 103, 104, 105}

	for _, id := range userIDs {
		wg.Add(1)
		ctx := context.WithValue(context.Background(), "UserID", id)
		go processRequest(ctx, &wg)
	}

	wg.Wait()
}

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	userID := ctx.Value("UserID").(int)
	fmt.Println("Processing request for User ID:", userID)
}
