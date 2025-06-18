package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main {
//Чтоб читать строку
	reader := bufio.NewReader(os.Stdin)
//Просим пользовотеля что-то написать
	fmt.Println("Напиши что-то")
//Читаем строку
	text, _ := reader.ReadString('\n')
//Делаем чтобы все хорошо было 
	text = string.ToLower(strings.TrimSpace(text))
//Разбиваем строку на слова 
	words := strings.Fields(text)
//Создаем хранилище всех слов
	wordCount := make(map[string]inp)
//цикл чтобы записывать все это в хранилище
	for _, word := range words {
		wordCount[word]++
	}
// Показываем результат
	fmt.Println("Результат посчета слов")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n")
	}
}