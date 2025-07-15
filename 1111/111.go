package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 8}

	fmt.Println(double(list))
}

func double(nums []int) []int {
	res := make([]int, len(nums))

	for i, num := range nums {
		res[i] = num * 2
	}
	return res
}
