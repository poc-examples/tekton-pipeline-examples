package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	// Print regular text
	fmt.Println("Hello, World!")

	// Print colored text
	color.Red("This is a red Hello World!")
	color.Green("This is a green Hello World!")
}