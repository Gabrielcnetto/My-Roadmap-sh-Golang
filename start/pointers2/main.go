package main

import "fmt"

func main() {
	i := 42
	var p *int
	p = &i
	fmt.Println(*p)
}
