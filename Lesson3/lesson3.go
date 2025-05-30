package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	var weight int
	var isHappy bool

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)

	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)

	fmt.Print("Введите ваш рост (в метрах): ")
	fmt.Scan(&height)

	fmt.Print("Введите ваш вес (в кг): ")
	fmt.Scan(&weight)

	fmt.Print("Вы счастливы? (true/false): ")
	fmt.Scan(&isHappy)

	fmt.Println("\nРезультаты:")
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Рост:", height)
	fmt.Println("Вec:", weight)
	fmt.Println("Счастлив?:", isHappy)

	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age < 60 {
		fmt.Println("Вы взрослый человек")
	} else {
		fmt.Println("Вы пенсионер.")
	}

	if isHappy {
		fmt.Println("Настроение отличное")
	} else {
		fmt.Println("Держитесь, все наладиться")
	}

	if height > 1.8 && weight <= 80 {
		fmt.Println("Ты в хорошей форме!")
	} else {
		fmt.Println("Над физической формой можно поработать :)")
	}
}
