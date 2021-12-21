package main

import "fmt"

func main() {
	biodata := map[string]string{
		"name":   "Lukman",
		"adress": "Lombok Tengah",
	}

	biodata["title"] = "Programmer"

	fmt.Println(biodata)
	//mengabil value dengan key
	fmt.Println(biodata["name"])
	fmt.Println(biodata["title"])
	fmt.Println(biodata["adress"])

	fmt.Println(len(biodata))

	var book = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "M Lukman Hakim"
	book["ups"] = "SAlah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
