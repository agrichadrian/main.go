package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var height float64
	var weight int
	var isHappy bool

	fmt.Println("Добро пожаловать в программу анализа состояния.")
	fmt.Println("Пожалуйста, введите следующую информацию.\n")

	// Имя
	fmt.Print("Введите своё имя: ")
	fmt.Scan(&name)

	// Возраст
	for {
		fmt.Print("Введите свой возраст: ")
		_, err := fmt.Scan(&age)
		if err == nil && age > 0 {
			break
		}
		fmt.Println("Ошибка: введите корректное положительное число.")
		fmt.Scanln()
	}

	// Рост
	for {
		fmt.Print("Введите свой рост в метрах (например, 1.82): ")
		_, err := fmt.Scan(&height)
		if err == nil && height > 0 {
			break
		}
		fmt.Println("Ошибка: введите корректное значение роста.")
		fmt.Scanln()
	}

	// Вес
	for {
		fmt.Print("Введите свой вес в кг: ")
		_, err := fmt.Scan(&weight)
		if err == nil && weight > 0 {
			break
		}
		fmt.Println("Ошибка: введите корректный вес.")
		fmt.Scanln()
	}

	// Настроение
	for {
		fmt.Print("Вы счастливы? (true/false): ")
		_, err := fmt.Scan(&isHappy)
		if err == nil {
			break
		}
		fmt.Println("Ошибка: введите true или false.")
		fmt.Scanln()
	}

	// Анализ
	fmt.Println("\nРезультаты анализа для", name, ":")

	fmt.Printf("Возраст: %d лет — ", age)
	if age < 18 {
		fmt.Println("вы молоды, у вас всё впереди.")
	} else if age < 60 {
		fmt.Println("вы в активном возрасте.")
	} else {
		fmt.Println("у вас большой жизненный опыт.")
	}

	fmt.Printf("Рост: %.2f м, Вес: %d кг — ", height, weight)
	if height > 1.8 && weight < 80 {
		fmt.Println("отличная физическая форма.")
	} else {
		fmt.Println("возможно, стоит уделить внимание здоровью.")
	}

	fmt.Print("Настроение: ")
	if isHappy {
		fmt.Println("отличное.")
	} else {
		fmt.Println("не переживайте, всё наладится.")
	}

	fmt.Println("\nСпасибо за использование программы.")
}
