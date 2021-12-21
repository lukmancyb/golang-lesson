package main

import "fmt"

func hello(name string) string {
	if name == "" {
		return "Gak dikenal"
	} else {
		return "Hello " + name
	}
}

func main() {

	result := hello("lukman")
	fmt.Println(result)

}
