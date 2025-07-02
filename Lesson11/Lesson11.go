package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal struct {
	Name   string
	Age    int
	Weight float64
	Type   string
}

func main() {
	var animals []Animal
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Сколько ты хочешь добавить животных? ")
	var count int
	fmt.Scan(&count)
	fmt.Scanln() // очистка после ввода числа

	for i := 0; i < count; i++ {
		fmt.Printf("\nЖивотное #%d\n", i+1)

		// Ввод имени
		var name string
		for {
			fmt.Print("Имя: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			name = strings.TrimSpace(text)
			if name == "" {
				fmt.Println("Ошибка: пустая строка. Повторите ввод.")
				continue
			}
			break
		}

		// Ввод типа
		var animalType string
		for {
			fmt.Print("Тип (например, собака, кошка): ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			animalType = strings.TrimSpace(text)
			if animalType == "" {
				fmt.Println("Ошибка: пустая строка. Повторите ввод.")
				continue
			}
			break
		}

		// Ввод возраста
		var age int
		for {
			fmt.Print("Возраст (целое число): ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			text = strings.TrimSpace(text)
			val, err := strconv.Atoi(text)
			if err != nil || val < 0 {
				fmt.Println("Ошибка: введите неотрицательное целое число.")
				continue
			}
			age = val
			break
		}

		// Ввод веса
		var weight float64
		for {
			fmt.Print("Вес (например, 4.5): ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			text = strings.TrimSpace(text)
			val, err := strconv.ParseFloat(text, 64)
			if err != nil || val < 0 {
				fmt.Println("Ошибка: введите неотрицательное число.")
				continue
			}
			weight = val
			break
		}

		// Добавление животного в список
		animal := Animal{
			Name:   name,
			Age:    age,
			Weight: weight,
			Type:   animalType,
		}
		animals = append(animals, animal)
	}

	// Вывод всех животных
	fmt.Println("\nВот список всех животных:")
	for i, a := range animals {
		fmt.Printf("\nЖивотное #%d:\n", i+1)
		fmt.Printf("Имя: %s\n", a.Name)
		fmt.Printf("Возраст: %d лет\n", a.Age)
		fmt.Printf("Вес: %.1f кг\n", a.Weight)
		fmt.Printf("Тип: %s\n", a.Type)
	}

	// Поиск животного по имени
	for {
		fmt.Print("\nХочешь найти животное по имени? (да/нет): ")
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		answer = strings.ToLower(strings.TrimSpace(answer))

		if answer == "нет" {
			// Просто выходим, если не нужно искать
			break
		} else if answer == "да" {
			fmt.Print("Введите имя для поиска: ")
			searchName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
			searchName = strings.TrimSpace(searchName)
			if searchName == "" {
				fmt.Println("Ошибка: пустая строка.")
				continue
			}

			// Ищем по списку
			found := false
			for _, a := range animals {
				if strings.EqualFold(a.Name, searchName) {
					fmt.Println("\nНайдено животное:")
					fmt.Printf("Имя: %s\n", a.Name)
					fmt.Printf("Возраст: %d лет\n", a.Age)
					fmt.Printf("Вес: %.1f кг\n", a.Weight)
					fmt.Printf("Тип: %s\n", a.Type)
					found = true
				}
			}
			if !found {
				fmt.Println("Животное с таким именем не найдено.")
			}
			break // после поиска выходим из цикла
		} else {
			fmt.Println("Пожалуйста, введи «да» или «нет».")
		}
	}
}
