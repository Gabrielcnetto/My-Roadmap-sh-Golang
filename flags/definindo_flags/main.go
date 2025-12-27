package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("name", "gabriel", "Nome do criador")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println(svar)
	fmt.Println(*wordPtr)

}
