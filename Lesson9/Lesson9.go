package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Print("Сколько чисел вы хотите ввести? ")
	_, err := fmt.Scan(&count)
	if err != nil || count <= 0 {
		fmt.Println("Ошибка ввода. Введите положительное число.")
		return
	}

	numbers := make([]int, count) // создаём срез нужной длины
	sum := 0
	max := 0

	fmt.Println("Введите числа:")

	for i := 0; i < count; i++ {
		fmt.Printf("%d: ", i+1)
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			i-- // вернёмся к текущей итерации
			continue
		}
		sum += numbers[i]
		if i == 0 || numbers[i] > max {
			max = numbers[i]
		}
	}

	fmt.Println("\nВведённые числа:")
	for i, num := range numbers {
		fmt.Printf("%d) %d\n", i+1, num)
	}

	fmt.Println("\nСумма всех чисел:", sum)
	fmt.Println("Максимальное число:", max)
}
