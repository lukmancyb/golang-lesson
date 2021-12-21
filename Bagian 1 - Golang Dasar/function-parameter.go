package main

import "fmt"

func hello(firstname string, lastname string) {
	fmt.Println("Hello world", firstname, lastname)
}

func main() {

	for i := 0; i < 10; i++ {
		hello("Lukman", "Hakim")
	}

}
