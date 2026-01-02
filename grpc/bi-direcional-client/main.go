package main

import (
	"context"
	"fmt"
	"io"
	"tidy/src/pb/shoppingcart"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	shoppingcart.ShoppingCartServiceServer
}

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Print("Erro ao criar conexão:", err.Error())
	}
	defer conn.Close()
	client := shoppingcart.NewShoppingCartServiceClient(conn)
	stream, err := client.AddItem(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}
	waitch := make(chan struct{})
	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(waitch)
				return
			}
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Received:", response)
		}
	}()
	items := []shoppingcart.AddProduct{
		shoppingcart.AddProduct{ProductId: 1, Quantity: 2, PriceUnit: 5},
		shoppingcart.AddProduct{ProductId: 5, Quantity: 3, PriceUnit: 2},
		shoppingcart.AddProduct{ProductId: 2, Quantity: 2, PriceUnit: 8},
	}
	for _, item := range items {
		if err := stream.Send(&item); err != nil {
			fmt.Println("Erro ao enviar:", err.Error())
		}
		fmt.Println("Send:", item)
	}
	stream.CloseSend()
	<-waitch
}
