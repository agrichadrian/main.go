package main

import "fmt"

// Функция умножения двух чисел
func multiply(a int, b int) int {
	return a * b
}

func main() {
	result := multiply(3, 5)
	fmt.Println("Произведение:", result)
}
