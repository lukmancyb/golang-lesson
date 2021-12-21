package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Lukman"
	names[1] = "Hakim"
	names[2] = "M"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var ages = [3]int{
		23,
		22,
		21,
	}

	fmt.Println(ages)

	fmt.Println(len(names))
	fmt.Println(len(ages))

}
