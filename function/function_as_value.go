package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Adil"))
}

func getGoodBye(nama string) string {
	return "Selamat Jalan " + nama + "!"
}
