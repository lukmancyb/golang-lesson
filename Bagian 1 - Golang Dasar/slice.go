package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"February",
		"maret",
		"april",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Juni Berubah"
	fmt.Println(slice1)
	slice1[0] = "Mei berubah"

	fmt.Println(months)
}
