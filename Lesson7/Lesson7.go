package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var fruits []string

	fmt.Println("Введи 3 своих любимых фрукта:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d. Фрукт: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fruits = append(fruits, input)
	}

	fmt.Println("\nТвои любимые фрукты:")
	for _, fruit := range fruits {
		fmt.Println("-", fruit)
	}

	fmt.Printf("Всего фруктов: %d\n", len(fruits))
}
