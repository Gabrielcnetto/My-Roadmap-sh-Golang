package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"tidy/src/pb/products"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	products.ProductServiceServer
}

func loadTLScred() (credentials.TransportCredentials, error) {
	serverCA, err := os.ReadFile("./src/cert/ca-cert.pem")
	if err != nil {
		fmt.Println(err.Error())
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(serverCA) {
		return nil, fmt.Errorf("Erro ao inicar certappend")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func main() {
	tlsData, err := loadTLScred()
	if err != nil {
		fmt.Println(err.Error())
	}
	conn, err := grpc.NewClient("0.0.0.0:9090", grpc.WithTransportCredentials(tlsData))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	client := products.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	productList, err := client.FindAll(ctx, &products.ListProductRequest{})
	if err != nil {
		fmt.Println("Erro:", err.Error())
	}
	fmt.Println(productList)
}
