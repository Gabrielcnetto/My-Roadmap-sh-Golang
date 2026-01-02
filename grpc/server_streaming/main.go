package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"tidy/src/pb/departament"

	"google.golang.org/grpc"
)

type server struct {
	departament.DepartamentServiceServer
}

func (s *server) ListPerson(req *departament.ListPersonRequest, srv departament.DepartamentService_ListPersonServer) error {
	file, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println("erro ao abrir:", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",") //aqui pega o content de cada linha que for passando
		id, _ := strconv.Atoi(data[0])
		name := data[1]
		email := data[2]
		income, _ := strconv.Atoi(data[3])
		departament_id, _ := strconv.Atoi(data[4])
		if int32(departament_id) == req.GetDepartamentId() {
			if err := srv.Send(&departament.ListPersonResponse{
				Id:            int32(id),
				Name:          name,
				Email:         email,
				Income:        int32(income),
				DepartamentId: int32(departament_id),
			}); err != nil {
				fmt.Println("Erro no send:", err.Error())
			}
		}
	}
	return nil
}
func main() {
	fmt.Println("iniciando server")
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()
	departament.RegisterDepartamentServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		fmt.Println(err.Error())
	}
}
