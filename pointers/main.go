package main

// * nas funcoes, pede o endereço daquele valor
/*
👉 Aqui você NÃO está pedindo o valor
👉 Você está pedindo o ENDEREÇO de um int
*/
func increment(value *int) {
	*value = *value * 10
}

type User struct {
	Name string
}

func (u *User) Setname(name string) {
	u.Name = name
}

// * pega o valor (desreferencia)
// & pega o endereço
// quando uma função pede *T, passamos &variável
// // dentro da função, usamos *param pra ler ou escrever o valor

func main() {
	//	i := 5
	/*
		 Passo a passo:
		i → valor 5
		&i → endereço de i
		esse endereço é passado para o parâmetro value
	*/
	//	increment(&i)
	//	client := &http.Client{}
	//	fmt.Println(i) // 20
	user := User{Name: "Gabriel"}
	user.Setname("Lucas")
}

var salario = 0

func SetNewSalario(new *int) {
	p := &salario
	*p = *new
}
