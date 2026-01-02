package main

import (
	"context"
	"fmt"
	"io"
	"tidy/src/pb/departament"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão:", err.Error())
	}
	defer conn.Close()
	client := departament.NewDepartamentServiceClient(conn)
	stream, err := client.ListPerson(context.Background(), &departament.ListPersonRequest{
		DepartamentId: 2,
	})
	if err != nil {
		fmt.Println("Erro ao usar stream:", err.Error())
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Erro no for:", err.Error())
		}
		fmt.Println("resposta:", response)
	}
}
