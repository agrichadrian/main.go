package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// safeYesNoInput — безопасный ввод: возвращает true для "да", false для "нет"
func safeYesNoInput(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		if input == "да" {
			return true
		} else if input == "нет" {
			return false
		} else {
			fmt.Println(" Введи только «да» или «нет». Попробуй ещё раз.")
		}
	}
}

func main() {
	isHappy := safeYesNoInput("Ты доволен? (да/нет): ")

	if isHappy {
		fmt.Println("Отлично!  Продолжай в том же духе.")
	} else {
		fmt.Println("Ничего страшного, всё обязательно улучшится! ")
	}
}
