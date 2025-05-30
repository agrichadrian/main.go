package main

import "fmt"

func main() {
	var name string = "Adrian"
	age := 24
	var height float64 = 1.87
	isHappy := true
	weight := 80

	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Рост:", height)
	fmt.Println("Счастлив?:", isHappy)
	fmt.Println("Вес:", weight)

	if age < 18 {
		fmt.Println("Вы несовершеннолетний.")
	} else if age < 60 {
		fmt.Println("Вы взрослый человек.")
	} else {
		fmt.Println("Вы пенсионер.")
	}

	if isHappy {
		fmt.Println("Настроение отличное!")
	} else {
		fmt.Println("Держись, всё наладится.")
	}

	if weight <= 80 {
		if height > 1.8 {
			fmt.Println("Ты в хорошей форме!")
		}
	} else {
		fmt.Println("Над физической формой можно поработать :)")
	}
}
