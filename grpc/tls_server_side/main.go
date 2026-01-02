package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"tidy/src/pb/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	products.ProductServiceServer
}

func LoadCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("./src/cert/server-cert.pem", "./src/cert/server-key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}

func (s *server) FindAll(ctx context.Context, req *products.ListProductRequest) (*products.ListProductResponse, error) {
	productList := make([]*products.Product, 0)
	productList = append(productList, &products.Product{Id: 1, Title: "Laptop"})
	return &products.ListProductResponse{Products: productList}, nil

}
func main() {
	fmt.Println("iniciando conexão")
	srv := server{}
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("Erro:", err.Error())
	}
	tlscredentials, err := LoadCredentials()
	if err != nil {
		fmt.Println(err.Error())
	}

	s := grpc.NewServer(
		grpc.Creds(tlscredentials),
	)
	products.RegisterProductServiceServer(s, &srv)
	if err := s.Serve(listener); err != nil {
		fmt.Println("Erro serve:", err.Error())
	}
}
