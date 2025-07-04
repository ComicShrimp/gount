package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	data, _ := os.ReadFile("./words.txt")
	fmt.Println(string(data))
}
