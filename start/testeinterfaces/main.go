package main

import (
	"fmt"
	"os"
)

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stdout, a...)
}

func main() {
	n := 10              // number
	s := "Hello!"        // string
	r := []rune("もしもし！") // slice of runes

	// Our final Println accepts different types of values
	// and go translate them in run time.
	Println(s, n, r)
}
