package main

import "fmt"

func hello() {
	fmt.Println("Hello world")
}

func main() {

	for i := 0; i < 10; i++ {
		hello()
	}

}
