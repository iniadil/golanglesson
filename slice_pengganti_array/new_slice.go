package main

import "fmt"

func main() {

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "adil"
	newSlice[1] = "pertama"
	//newSlice[2] = "ibnu" // error dan harusnya menggunakan append karena diluar len tapi masih dalam kapasitas (capcity
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}
