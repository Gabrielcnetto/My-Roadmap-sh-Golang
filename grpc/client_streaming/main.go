package main

import (
	"fmt"
	"io"
	"net"
	"tidy/src/pb/calc"

	"google.golang.org/grpc"
)

type server struct {
	calc.CalcServiceServer
}

func (s *server) Calc(stream calc.CalcService_CalcServer) error {
	var quantity int32 = 0
	var _total int32 = 0
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calc.Output{
				Quantity: quantity,
				Total:    float32(_total),
				Average:  (float32(_total) / float32(quantity)),
			})
		}
		quantity++
		_total = input.GetValue()
	}
}

func main() {
	fmt.Println("Server iniciado")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Erro ao iniciar net", err)
	}
	s := grpc.NewServer()
	calc.RegisterCalcServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Println("Erro #2", err)
	}
}
