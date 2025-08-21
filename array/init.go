package main

import "fmt"

func main() {
	var names = [3]string{
		"adil",
		"muhammad",
		"Muhammad", // wajib memiliki koma walaupun di akhir value
	}

	fmt.Println(names)
}
