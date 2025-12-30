package main

import (
	"context"
	"fmt"
	"tidy/src/pb/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Errorf("Erro ao iniciar:", err.Error())
	}
	defer conn.Close()
	FindAll(conn)
	newP := products.Product{
		Name:        "Nome teste",
		Description: "nome de teste para usar",
		Price:       34.90,
		Quantity:    7,
	}
	Insert(conn, &newP)
	fmt.Println("___________Buscando Novamente___________")
	FindAll(conn)
}
func FindAll(conn *grpc.ClientConn) {
	client := products.NewProductServiceClient(conn)
	list, err := client.FindAll(context.Background(), &products.Product{})
	if err != nil {
		fmt.Errorf("Erro ao pegar a lista:", err.Error())
	}
	fmt.Printf("Products: %+v\n", list)
}

func Insert(conn *grpc.ClientConn, item *products.Product) {
	client := products.NewProductServiceClient(conn)
	response, err := client.Create(context.Background(), item)
	if err != nil {
		fmt.Errorf("Erro para adicionar:", err.Error())
	}
	fmt.Println("Item adicionado com sucesso:", response)
}
