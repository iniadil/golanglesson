package main

import "fmt"

func main() {
	// len menghitung isi array
	var names = [...]string{
		"adil",
		"muhammad",
	}

	fmt.Println(len(names))

}
