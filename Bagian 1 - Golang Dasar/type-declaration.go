package main

import "fmt"

func main() {
	type NoKtp = string
	type status = bool

	var noKtpLukman NoKtp = "89809090"
	var statusPerkawinan status = true

	fmt.Println(noKtpLukman)
	fmt.Println(statusPerkawinan)
}
