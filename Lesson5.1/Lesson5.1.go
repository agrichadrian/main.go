package main

import (
	"fmt"
	"strings"
)

func main() {
	var sleepHours float64
	var input string
	var didExercise, didShower, hadBreakfast bool

	fmt.Println("Привет! Давай оценим, насколько продуктивно начался твой день.")

	fmt.Print("Сколько часов ты сегодня спал ?")
	fmt.Scan(&sleepHours)

	fmt.Print("Ты делал зарядку? (да/нет): ")
	fmt.Scan(&input)
	didExercise = strings.ToLower(input) == "да"

	fmt.Print("Ты принял душ? (да/нет): ")
	fmt.Scan(&input)
	didShower = strings.ToLower(input) == "да"

	fmt.Print("Ты позавтракал? (да/нет): ")
	fmt.Scan(&input)
	hadBreakfast = strings.ToLower(input) == "да"

	score := 0

	if sleepHours >= 7 {
		score++
	}
	if didExercise {
		score++
	}
	if didShower {
		score++
	}
	if hadBreakfast {
		score++
	}

	fmt.Print("\n Результат анализа")
	switch score {
	case 4:
		fmt.Println("Ты заряжен на 100%! Отличное начало!")
	case 3:
		fmt.Println("Хорошее утро! Ещё чуть-чуть — и ты на пике!")
	case 2:
		fmt.Println("Средний старт — можно усилиться завтра.")
	case 1:
		fmt.Println("Утро было тяжёлым. Но ты справился — это главное.")
	default:
		fmt.Println("Сегодня было сложно, но ты молодец, что продолжаешь.")
	}

	fmt.Println("Продолжай в том же духе — даже маленькие шаги важны!")
}
