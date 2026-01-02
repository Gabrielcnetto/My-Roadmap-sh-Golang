package main

import (
	"fmt"
	"io"
	"net"
	"tidy/src/pb/shoppingcart"

	"google.golang.org/grpc"
)

type server struct {
	shoppingcart.ShoppingCartServiceServer
}

func (s *server) AddItem(srv shoppingcart.ShoppingCartService_AddItemServer) error {
	for {
		quantity_items := 0
		priceTotal := 0.0
		newitem, err := srv.Recv()
		if err == io.EOF {
			return srv.Send(&shoppingcart.ShoppingCartTotal{
				Quantity: int32(quantity_items),
				Amount:   priceTotal,
			})
		}
		if err != nil {
			fmt.Println("Erro on receive:", err.Error())
		}
		quantity_items += int(newitem.GetQuantity())
		priceTotal += float64((newitem.PriceUnit * float64(newitem.GetQuantity())))
		if err := srv.Send(&shoppingcart.ShoppingCartTotal{
			Quantity: int32(quantity_items),
			Amount:   priceTotal,
		}); err != nil {
			fmt.Println("Erro ao enviar:", err.Error())
		}
	}
}

func main() {
	fmt.Println("Iniciando servidor")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	shoppingcart.RegisterShoppingCartServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Println(err.Error())
	}
}
