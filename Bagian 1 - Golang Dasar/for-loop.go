package main

import "fmt"

func main() {

	//For loop
	// for counter := 1; counter < 10; counter++ {
	// 	fmt.Println("Perulangan", counter)
	// }

	slice := []string{"eko", "kurniawan", "khanedy"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println("index", i, "=", slice[i])
	// }

	// for i, value := range slice {
	// 	fmt.Println("Index ke", i, "adalah ", value)
	// }

	for _, value := range slice {
		fmt.Println("adalah ", value)
	}

	//map
	person := make(map[string]string)

	person["title"] = "Programmer"
	person["name"] = "Lukman"

	for key, value := range person {
		fmt.Println("Keynya", key, "valuenya", value)
	}

}
