package main

import (
	"fmt"
	"os"
	"tidy/src/pb/users"

	"google.golang.org/protobuf/proto"
)

func CreateData() {
	user1 := users.User{
		Id:       1234,
		Name:     "Gabriel",
		Email:    "Gabrielcarlosnettoo@gmail.com",
		Password: "acesso",
	}
	out, err := proto.Marshal(&user1)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	errFile := os.WriteFile("dados.txt", out, 0644)
	if errFile != nil {
		fmt.Errorf(errFile.Error())
	}
}
func ReadData() {
	var user1 users.User
	dados, err := os.ReadFile("dados.txt")
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	if err := proto.Unmarshal(dados, &user1); err != nil {
		fmt.Println("Error2:", err.Error())
	}
	fmt.Print("user:", user1)

}
func main() {
	ReadData()

}
