package main

import "fmt"

func main() {
	var benar bool = true
	var guest bool = true

	var hasil bool = benar && guest
	fmt.Println(hasil)

	var hasil2 bool = benar || guest
	fmt.Println(hasil2)

	var hasil3 bool = !benar
	fmt.Println(hasil3)
}
