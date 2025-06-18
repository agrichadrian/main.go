package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Делаем так чтобы программа умела читать ввод
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите любой текст")

	//Делаем чбы она могла читать
	text, _ := reader.ReadString('\n')
	//Делаем чтобы она читала без пробелов и заглавных
	text = strings.ToLower(strings.TrimSpace(text))

	//Удаляем знаки
	replacer := strings.NewReplacer(
		".", "",
		",", "",
		"!", "",
		"?", "",
		":", "",
		";", "",
		"-", "",
		"\"", "",
		"'", "",
		"(", "",
		")", "",
		"—", "",
	)
	cleanText := replacer.Replace(text)
	// Разбиваем строку на отдельные слова
	words := strings.Fields(cleanText)
	// Слова которые будем игнорировать
	stopWords := map[string]bool{
		"и": true, "в": true, "на": true, "это": true, "а": true,
		"по": true, "с": true, "что": true, "как": true, "но": true,
		"у": true, "за": true, "от": true, "до": true, "из": true,
	}
	// Создаем пустую map для храненияч количества повторений
	wordCount := make(map[string]int)
	// Проходим по каждому слову
	for _, word := range words {
		if stopWords[word] {
			continue
		}
		word = strings.ToLower(word)
		wordCount[word]++ //Увиличиваем сотчик для этого слова
	}
	// Вывод результата
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
	//Поиск самого частого слова
	var mostFrequentWord string
	maxCount := 0

	for word, count := range wordCount {
		if count > maxCount {
			mostFrequentWord = word
			maxCount = count
		}
	}
	if mostFrequentWord != "" {
		fmt.Printf("\nСамое частое слово: \"%s\" (встречается %d раз)\n", mostFrequentWord, maxCount)
	} else {
		fmt.Println("\nНет слов для анализа.")
	}
}
