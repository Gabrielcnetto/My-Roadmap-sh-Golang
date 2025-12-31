package main

import (
	"context"
	"fmt"
	"tidy/src/pb/calc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("erro conn:", err)
	}
	defer conn.Close()
	client := calc.NewCalcServiceClient(conn)
	stream, err := client.Calc(context.Background())
	if err != nil {
		fmt.Println("Erro ao usar stream:", err.Error())
	}
	nums := []int32{1, 2, 3, 5}
	for _, item := range nums {
		if err := stream.Send(&calc.Input{Value: item}); err != nil {
			fmt.Println("Erro ao send:", err.Error())
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("Erro ao close:", err.Error())
	}
	fmt.Println("response final:", response)
}
