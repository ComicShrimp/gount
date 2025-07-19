package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	_ = data

	wordCount := 1

	for _, word := range data {
		if word == ' ' {
			wordCount++
			fmt.Println("Space detected")
		}
	}

	fmt.Println(wordCount)
}
