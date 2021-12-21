package main

import "fmt"

func main() {
	name := "Luk"
	switch name {
	case "Lukman":
		fmt.Println("Hello lukman")
	case "Hakim":
		fmt.Println("Hello Hakim")
	default:
		fmt.Println("GAK DIKENAL")
	}

	//Sort statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Kepanjangan")
	case false:
		fmt.Println("Benar sudah")
	}

	//Switch tanpa ekspresi

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Kepanjangan")
	case length > 10:
		fmt.Println("Lumayan panjang")
	default:
		fmt.Println("Benar")
	}
}
