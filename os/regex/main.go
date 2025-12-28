package main

import (
	"fmt"
	"regexp"
)

func main() {
	inputText := "I love new york city"
	match, err := regexp.MatchString("[A-z]ork", inputText)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}
}
