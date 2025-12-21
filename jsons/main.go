package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome   string `json:"pessoa"`
	Idade  int    `json:"idade,omitempty"` //quando usa o omitempty ele nao vai aceitar nulos ou zero
	Cidade string `json:"-"`               //aqui a gente exclui isso do json
}

/*
func main() {
	_pessoa := Pessoa{
		Nome:  "Gabriel",
		Idade: 21,
	}
	data, err := json.MarshalIndent(_pessoa, "", "")
	if err != nil {
		fmt.Println("Erro ao converter:", err.Error())
	}
	fmt.Println(string(data))
}
*/

func main() {
	//criando o json
	var raw Pessoa
	//jsonData := []byte(fmt.Sprintf(`{"nome":"Gabriel","idade":%v}`, 0))
	jsonData := []byte(`{"nome":"Gabriel","idade":}`)
	if err := json.Unmarshal(jsonData, &raw); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsonData))
}
