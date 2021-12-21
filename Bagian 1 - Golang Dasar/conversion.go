package main

import "fmt"

func main() {
	var nilai32 int32 = 23789
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	//kembali ke nilai awal jika tipe data lebih kecil
	fmt.Println(nilai8)

	var nama = "Lukman"
	var e = nama[0]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(eString)

}
