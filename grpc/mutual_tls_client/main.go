package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"tidy/src/pb/products"

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
	clientCert, err := tls.LoadX509KeyPair("./src/cert/client-cert.pem", "./src/cert/client-key.pem")
	if err != nil {
		fmt.Println(err.Error())
	}
	config := &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
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
	productList, err := client.FindAll(context.Background(), &products.ListProductRequest{})
	if err != nil {
		fmt.Println("Erro:", err.Error())
	}
	fmt.Println(productList)
}
