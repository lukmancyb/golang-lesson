package main

import "fmt"

func getName() (firstname string, lastname string) {
	firstname = "lukman"
	lastname = "hakim"
	return
}

func main() {

	a, b := getName()
	fmt.Println(a, b)

}
