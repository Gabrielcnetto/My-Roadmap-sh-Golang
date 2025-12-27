package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "world", "nome a ser exibido")
	debug := flag.Bool("debug", false, "ativa modo debug")

	flag.Parse()

	fmt.Println("Hello,", *name)

	if *debug {
		fmt.Println("Debug ativo")
	}
}
