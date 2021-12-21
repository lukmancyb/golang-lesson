package main

import "fmt"

func main() {
	var nama string
	nama = "lukman"
	fmt.Println(nama)

	nama = "Hakim"
	fmt.Println(nama)

	//var bisa diInisialisasi langsung
	var age = 30
	fmt.Println(age)

	// Iisialisasi variabel tanpa kata var ->syarat harus diberi nilai langsung
	pekerjaan := "programmer"
	fmt.Println(pekerjaan)

	//Multiple inisialisai

	var (
		umur          = 21
		jenis_kelamin = "Pria"
	)

	fmt.Println(umur)
	fmt.Println(jenis_kelamin)

}
