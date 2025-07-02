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
			// 2.3) Возвращаю то что накалякал в файл
			line := scanner.Text()
			// 2.4) Разделяю строку на две части
			parts := strings.SplitN(line, "|", 2)
			// 2.5) Присваюваю каждой части свою функцию
			if len(parts) == 2 {
				name := parts[0]
				phone := parts[1]
				// 2.6) Записываю в пам, присвамваю номер к имени
				phoonebook[name] = phone
			}
		}
		// 2.6) Закрываю
		file.Close()
	}
	// 3) Апать создаю сканер строки которой мы написали
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
		// 4.2) Создаем переменную инт для чтения
		var choice int
		// 4.3) Читаем что ввели
		_, err := fmt.Scan(&choice)
		fmt.Scanln()
		// 4.3.1) Обрабатываем критинизм юзера
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз")
			// 4.3.2) Очищаем калаяки пользовотеля
			scanner.Scan()
			continue
		}
		// 5) Создаем свич систэм для нашей переменной
		switch choice {
		//если юзер вводит 1
		case 1:
			// 6.1) Создаем условие если длина равнаяется 0
			if len(phoonebook) == 0 {
				fmt.Println("Телефонная книга пуста.")
				// 6.1.1 Иначе выводим цикл по ранжэ нашей телефоннкниги
			} else {
				fmt.Println("Список контактов:")
				// 6.1.2 Ну тут все понятно, прин всей нашей книжки наме и номера
				for name, phone := range phoonebook {
					fmt.Printf("%s: %s\n", name, phone)
				}
			}
			//если юзер вводит 2
		case 2:
			// 7 выводит текст
			fmt.Print("Введите имя: ")
			// 7.1 ты вводишь текста
			scanner.Scan()
			// 7.2 возвращает текст который ты вел
			name := scanner.Text()
			// 7.3 условие которое сравнивает написанное нами с тем что есть в книге
			if phone, ok := phoonebook[name]; ok {
				fmt.Printf("Номер для %s: %s\n", name, phone)
			}
			//если юзер вводит 3
		case 3:
			// 8 вводилка имени
			fmt.Println("Имя нового контакта: ")
			scanner.Scan()
			name := strings.TrimSpace(scanner.Text())
			if name == "" {
				fmt.Println("Ошибка: имя не может быть пустым.")
				continue
			}
			// 8.1 вводилка номера
			fmt.Println("\nНомер телефона: ")
			scanner.Scan()
			phone := strings.TrimSpace(scanner.Text())
			if phone == "" {
				fmt.Println("Ошибка: номер не может быть пустым.")
				continue
			}
			// 8.2 наши вводилки отправляються в наш phonebook мэп
			phoonebook[name] = phone
			fmt.Println("Контакт добавлен")
			//если юзер вводит 4
		case 4:
			// 9 ввод именя
			fmt.Print("Кого удалить? Введите имя: ")
			scanner.Scan()
			name := scanner.Text()
			// 9.1 условие при совпадание с именем из нашего мапа
			if _, exsist := phoonebook[name]; exsist {
				// 9.1.1 удаляет его
				delete(phoonebook, name)
				fmt.Println(("Контакт удален."))
				// 9.1.2 не удалает его
			} else {
				fmt.Println("Такого контакта нет.")
			}
		case 5:
			file, err := os.Create("phonebook.txt")
			if err != nil {
				fmt.Println("Не удалось сохранить файл.")
			} else {
				defer file.Close()
				for name, phone := range phoonebook {
					file.WriteString(name + "|" + phone + "\n")
				}
			}
			fmt.Println("Кеина успешно сохранена")
		case 6:
			fmt.Println("Выход. Довстречи")
			return

		default:
			fmt.Println("Неверный выбор.")
		}
	}
}
