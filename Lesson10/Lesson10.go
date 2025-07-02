package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите текст:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Println("Ошибка: пустой ввод.")
		return
	}

	words := strings.Fields(input)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\nСтатистика слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
