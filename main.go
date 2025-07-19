package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := countWords(data)

	fmt.Println(wordCount)
}

func countWords(data []byte) int {
	wordCount := 1

	for _, word := range data {
		if word == ' ' {
			wordCount++
		}
	}

	return wordCount
}
