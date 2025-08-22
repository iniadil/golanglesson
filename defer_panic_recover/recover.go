package main

import "fmt"

func finish() {
	fmt.Println("finish")
	messageError := recover()
	fmt.Println("Error pada: ", messageError)
}

func initAppRecover(error bool) {
	defer finish()

	if error {
		panic("ada erorrrrrrr")
	} else {
		fmt.Println(" aplikasi berhasil jalan")
	}

	// kode dibawah adalah salah karena tidak akan di eksekusi
	// recover harusnya di buat pada function defer agar tereksekusi
	// messageError := recover()
	// fmt.Println("Error pada: ", messageError)
}

func main() {
	initAppRecover(true)
	fmt.Println("Muhammad Adil")
}
