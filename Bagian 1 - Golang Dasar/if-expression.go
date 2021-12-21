package main

import "fmt"

func main() {
	var name = "eokllklk"

	if name == "eko" {
		fmt.Println("Hello eko")
	} else if name == "eo" {
		fmt.Println("Hello eo")
	} else {
		fmt.Println("Kenalan dong")
	}

	if panjang := len(name); panjang > 5 {
		fmt.Println("Kepanjangan cuy")
	} else {
		fmt.Println("Sudah benar")
	}

}
