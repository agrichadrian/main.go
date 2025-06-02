package main

import "fmt"

// Функция умножения двух чисел
func multiply(a int, b int) int {
	return a * b
}

// Функция расчёта ИМТ
func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func main() {
	// Умножение
	result := multiply(3, 5)
	fmt.Println("Произведение:", result)

	// Ввод веса и роста
	var weight float64
	var height float64

	for {
		fmt.Print("Введите ваш вес (кг): ")
		_, err := fmt.Scan(&weight)
		if err == nil && weight > 0 {
			break
		}
		fmt.Println("Введите корректный вес.")
		fmt.Scanln()
	}

	for {
		fmt.Print("Введите ваш рост (в метрах, например 1.75): ")
		_, err := fmt.Scan(&height)
		if err == nil && height > 0 {
			break
		}
		fmt.Println("Введите корректный рост.")
		fmt.Scanln()
	}

	// Расчёт и вывод ИМТ
	bmi := calculateBMI(weight, height)
	fmt.Printf("Ваш ИМТ: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("Недостаточный вес.")
	} else if bmi < 25 {
		fmt.Println("Нормальный вес.")
	} else if bmi < 30 {
		fmt.Println("Избыточный вес.")
	} else {
		fmt.Println("Ожирение.")
	}
}
