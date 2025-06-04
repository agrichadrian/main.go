package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var books []string

	fmt.Println("Привет! Давай я запишу твои любимые книги.")

	// Цикл ввода
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d. Книга: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		books = append(books, input)
	}

	// Вывод результата
	fmt.Println("\nВот список твоих любимых книг:")
	for i, book := range books {
		fmt.Printf("%d) %s\n", i+1, book)
	}

	fmt.Printf("\nВсего книг: %d\n", len(books))
}
