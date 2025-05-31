package main

import "fmt"

func main() {
	var name string
	var age int
	var weight int
	var health bool
	var tipe string

	// Имя
	for {
		fmt.Print("Введите кличку животного: ")
		_, err := fmt.Scan(&name)
		if err == nil && name != "" {
			break
		}
		fmt.Println("Пожалуйста, введите корректное имя.")
		fmt.Scanln()
	}

	// Возраст
	for {
		fmt.Print("Введите возраст животного: ")
		_, err := fmt.Scan(&age)
		if err == nil && age > 0 {
			break
		}
		fmt.Println("Пожалуйста, введите корректный возраст.")
		fmt.Scanln()
	}

	// Вес
	for {
		fmt.Print("Введите вес животного: ")
		_, err := fmt.Scan(&weight)
		if err == nil && weight > 0 {
			break
		}
		fmt.Println("Пожалуйста, введите корректный вес.")
		fmt.Scanln()
	}

	// Тип
	for {
		fmt.Print("Введите тип животного (например: Кошка, Собака): ")
		_, err := fmt.Scan(&tipe)
		if err == nil && tipe != "" {
			break
		}
		fmt.Println("Пожалуйста, введите корректный тип.")
		fmt.Scanln()
	}

	// Здоровье
	for {
		fmt.Print("Здоров ли питомец? (true/false): ")
		_, err := fmt.Scan(&health)
		if err == nil {
			break
		}
		fmt.Println("Пожалуйста, введите true или false.")
		fmt.Scanln()
	}

	// Анализ
	fmt.Println("\nРезультаты анализа:")
	fmt.Println("Имя питомца:", name)
	fmt.Println("Возраст питомца:", age)

	if age < 3 {
		fmt.Println("Питомец очень молодой.")
	} else if age < 10 {
		fmt.Println("Питомец взрослый.")
	} else {
		fmt.Println("Питомец пожилой.")
	}

	fmt.Println("Вес питомца:", weight)

	fmt.Print("Комментарий: ")
	switch tipe {
	case "Собака":
		fmt.Println("Собаки обожают гулять.")
	case "Кошка":
		fmt.Println("Кошки любят спать.")
	case "Попугай":
		fmt.Println("Попугаи любят летать.")
	default:
		fmt.Println("Это интересный питомец!")
	}

	fmt.Print("Здоровье: ")
	if health {
		fmt.Println("всё в порядке.")
	} else {
		fmt.Println("требуется внимание.")
	}
}
