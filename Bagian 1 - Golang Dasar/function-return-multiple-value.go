package main

import "fmt"

func getName() (string, string) {
	return "lukman", "hakim"
}

func main() {

	fname, _ := getName()
	fmt.Println(fname)

}
