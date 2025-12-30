package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"tidy/src/pb/products"
	"tidy/src/repository"

	"google.golang.org/grpc"
)

type server struct {
	products.ProductServiceServer
	prodRepo repository.ProductRepository
}

func (s *server) Create(ctx context.Context, product *products.Product) (*products.Product, error) {
	newproduct, err := s.prodRepo.Create(*product)
	if err != nil {
		return product, err
	}
	return &newproduct, nil

}
func (s *server) FindAll(ctx context.Context, product *products.Product) (*products.ProductList, error) {
	items, err := s.prodRepo.FindAll()
	return &items, err
}

func main() {
	fmt.Println("Iniciando servidor")
	srv := server{
		prodRepo: repository.ProductRepository{},
	}
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	products.RegisterProductServiceServer(s, &srv)
	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
