package main

import (
	"flag"
	"fmt"
)

func main() {
	channel := flag.String("channelName", "Gabriel", "Nome do cnaal")
	flag.Parse()
	fmt.Println(*channel)
}
