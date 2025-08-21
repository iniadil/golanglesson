package main

import "fmt"

func main() {
	names := [...]string{
		"adil",
		"javascript",
		"ibnu",
	}

	// buat dulu variable singkatnya; lalu if statementnya
	// ini berfungsi untuk mempersingkat statement agar mempermudah
	if panjang := len(names); panjang <= 2 {
		fmt.Println("benar")
	} else {
		fmt.Println("terlalu banyak")
	}
}
