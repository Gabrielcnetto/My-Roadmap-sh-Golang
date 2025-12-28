package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	command := flag.String("command", "", "Command to execute")
	flag.Parse()

	if *command == "" {
		fmt.Println("Please provide a command.")
		flag.Usage()
		return
	}

	switch *command {
	case "greet":
		fmt.Println("Hello, World!")
	case "time":
		fmt.Println("Current time:", time.Now().Format("15:04:05"))
	default:
		fmt.Println("Unknown command:", command)
	}
}
