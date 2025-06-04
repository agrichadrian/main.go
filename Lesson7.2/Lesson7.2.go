package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1) Ввод количества фильмов (1–10)
	var count int
	for {
		fmt.Print("Введите число фильмов (1–10): ")
		_, err := fmt.Scan(&count)
		fmt.Scanln() // очищаем остаток ввода
		if err == nil && count >= 1 && count <= 10 {
			break
		}
		fmt.Println("Нужно число от 1 до 10. Попробуйте снова.")
	}

	// Объявляем срез для хранения названий фильмов
	var movies []string

	// 2) Ввод самих названий - count раз
	for i := 1; i <= count; i++ {
		for {
			fmt.Printf("%d. Название фильма: ", i)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("Название не может быть пустым. Попробуйте снова.")
				continue
			}
			movies = append(movies, input)
			break
		}
	}

	// 3) Сохранение в файл (необязательно, но оставляем)
	file, err := os.Create("movies_list.txt")
	if err == nil {
		defer file.Close()
		for _, movie := range movies {
			file.WriteString(movie + "\n")
		}
	} // Если ошибка — просто пропускаем запись

	// 4) Вывод первоначального списка
	fmt.Println("\nТекущий список фильмов:")
	for i, movie := range movies {
		fmt.Printf("%d) %s\n", i+1, movie)
	}

	// 5) Редактирование элементов списка
	for {
		fmt.Print("\nХотите отредактировать название фильма? (да/нет): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "нет" {
			break // выходим из цикла редактирования
		} else if answer == "да" {
			// Выводим список заново, чтобы пользователь видел индексы
			fmt.Println("\nСписок фильмов:")
			for i, movie := range movies {
				fmt.Printf("%d) %s\n", i+1, movie)
			}

			// Спрашиваем номер
			fmt.Print("Введите номер фильма для изменения: ")
			var index int
			_, err := fmt.Scan(&index)
			fmt.Scanln() // очищаем буфер
			if err != nil || index < 1 || index > len(movies) {
				fmt.Println("Некорректный номер. Попробуйте ещё раз.")
				continue
			}

			// Ввод нового названия
			fmt.Print("Введите новое название: ")
			newTitle, _ := reader.ReadString('\n')
			newTitle = strings.TrimSpace(newTitle)
			if newTitle == "" {
				fmt.Println("Название не может быть пустым.")
				continue
			}

			// Обновляем значение в срезе
			movies[index-1] = newTitle
			fmt.Println("Название фильма обновлено.")
		} else {
			fmt.Println("Нужно ввести «да» или «нет».")
		}
	}

	// После редактирования выводим текущий список
	fmt.Println("\nОбновлённый список фильмов:")
	for i, movie := range movies {
		fmt.Printf("%d) %s\n", i+1, movie)
	}

	// 6) Удаление элементов из списка
	for {
		fmt.Print("\nХотите удалить фильм из списка? (да/нет): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "нет" {
			break // выходим, если не удаляем
		} else if answer == "да" {
			// Показываем текущий список
			fmt.Println("\nСписок фильмов:")
			for i, movie := range movies {
				fmt.Printf("%d) %s\n", i+1, movie)
			}

			// Спрашиваем номер для удаления
			fmt.Print("Введите номер фильма для удаления: ")
			var index int
			_, err := fmt.Scan(&index)
			fmt.Scanln() // очищаем буфер
			if err != nil || index < 1 || index > len(movies) {
				fmt.Println("Некорректный номер. Попробуйте ещё раз.")
				continue
			}

			// Удаляем элемент из среза: создаём новый срез без этого элемента
			movies = append(movies[:index-1], movies[index:]...)
			fmt.Println("Фильм удалён из списка.")
		} else {
			fmt.Println("Нужно ввести «да» или «нет».")
		}
	}

	// После удаления выводим текущий список
	fmt.Println("\nСписок фильмов после удаления:")
	for i, movie := range movies {
		fmt.Printf("%d) %s\n", i+1, movie)
	}

	// 7) Поиск по названию
	for {
		fmt.Print("\nХотите найти фильм по названию? (да/нет): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(strings.ToLower(answer))

		if answer == "нет" {
			break // выходим, если не ищем
		} else if answer == "да" {
			// Запрашиваем ключевое слово для поиска
			fmt.Print("Введите слово для поиска (без учёта регистра): ")
			query, _ := reader.ReadString('\n')
			query = strings.TrimSpace(strings.ToLower(query))

			if query == "" {
				fmt.Println("Пустая строка. Попробуйте снова.")
				continue
			}

			// Перебираем все фильмы и ищем совпадения
			found := false
			fmt.Println("\nРезультаты поиска:")
			for i, movie := range movies {
				if strings.Contains(strings.ToLower(movie), query) {
					fmt.Printf("%d) %s\n", i+1, movie)
					found = true
				}
			}
			if !found {
				fmt.Println("Ничего не найдено по запросу.")
			}
		} else {
			fmt.Println("Нужно ввести «да» или «нет».")
		}
	}

	// 8) Финальный вывод итогового списка
	fmt.Println("\nИтоговый список фильмов:")
	for i, movie := range movies {
		fmt.Printf("%d) %s\n", i+1, movie)
	}

	fmt.Println("\nРабота завершена.")
}
