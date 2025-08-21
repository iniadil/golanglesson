package main

import "fmt"

func main() {
	var age int16 = 17
	description := checkUmur(age)
	fmt.Println(description)
}

// parameter tidak wajib
func checkUmur(age int16) string {
	switch {
	case age > 18:
		return "terlalu tua"
	case age < 18:
		return "masih muda"
	default:
		return "tidak keduanya"
	}
}
