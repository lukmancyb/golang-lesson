package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	data := sumAll(10, 11, 12, 13)
	fmt.Println(data)

	slice := []int{10, 20, 30, 40}
	total := sumAll(slice...)
	fmt.Println(total)
}
