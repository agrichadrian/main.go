package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1) СОздаем мап для телефонной книги
	phoonebook := make(map[string]string)

	// 2) Загруска из файла
	file, err := os.Open("phonebook.txt")
	if err == nil {
		// 2.1) Создаю сканер чтобы читать файл по строчно
		scanner := bufio.NewScanner(file)
		// 2.2) Юзаю сканер
		for scanner.Scan() {
			// 2.3) Возвращаю то что накалякал
			line := scanner.Text()
			// 2.4) Разделяю строку на две части
			parts := strings.SplitN(line, "|", 2)
			// 2.5) Присваюваю каждой части свою функцию
			if len(parts) == 2 {
				name := parts[0]
				phone := parts[1]
				// 2.6) Записываю в пам, примваюваю номер к имени
				phoonebook[name] = phone
			}
		}
		// 2.6) Закрываю
		file.Close()
	}
	// 3) Апать создаю сканер
	scanner := bufio.NewScanner(os.Stdin)
	// 4) Запускаем цикл
	for {
		// 4.1) Выводим функционал нашего цикла
		fmt.Println("\nЧто ты хочешь сделать?")
		fmt.Println("1. Показать всю книгу")
		fmt.Println("2. Найти номер по имени")
		fmt.Println("3. Добавить новую запись")
		fmt.Println("4. Удалить запись")
		fmt.Println("5. Сохранить фаил")
		fmt.Println("6. Выйти")
		fmt.Print("Выбор: ")
		// 4.2) Создаем переменную инт
		var choice int
		// 4.3) Читаем что ввели
		_, err := fmt.Scan(&choice)
		// 4.3.1) Обрабатываем критинизм юзера
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз")
			// 4.3.2) Очищаем калаяки пользовотеля
			scanner.Scan()
			continue
		}
		switch choice {
		case 1:
			if len(phoonebook) == 0 {
				fmt.Println("Телефонная книга пуста.")
			} else {
				fmt.Println("Список контактов:")
				for name, phone := range phoonebook {
					fmt.Printf("%s: %s\n", name, phone)
				}
			}
		case 2:
			fmt.Print("Введите имя: ")
			scanner.Scan()
		}
	}
}
