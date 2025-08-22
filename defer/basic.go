package main

import "fmt"

func success() {
	fmt.Println("Success")
}

func initApp() {
	// ini akan dipanggil setelah semua eksekusi selesai
	// cocok untuk pemanggilan success atau lainnya
	defer success()

	a := 1
	b := 2
	c := a * b
	fmt.Println(c)
}

func main() {
	initApp()
}
