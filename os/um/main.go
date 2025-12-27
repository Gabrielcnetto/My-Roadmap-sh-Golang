package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
	LENDO UM ARQUIVO TXT

	func main() {
		file, err := os.Open("../golang.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		buf := make([]byte, 1024)
		n, err := file.Read(buf)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))

}
*/
/*criando pastas
func main() {
	err := os.Mkdir("example_dir", 0755)
	if err != nil {
		fmt.Println("Erro:", err.Error())
		return
	}
	err = os.MkdirAll("example_dir/subdir/subsubdir", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
*/
/*lendo com bufior, pra ler parte por parte e nao carreagar tudo
func main() {
	file, err := os.Open("./text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
*/

func main() {
	file, err := os.Create("teste.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	log.Writer()
	writer.WriteString("package main\nfunc teste(){}")
	writer.Flush()

}
