package main

import "fmt"

func main() {
	var a = 1
	if a == 1 {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	if a != 1 {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	var bool bool = a == 1
	fmt.Println(bool)
}
