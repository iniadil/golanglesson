package main

import "fmt"

func success() {
	fmt.Println("Kode Selesai")
}

func initApp(error bool) {
	// cek defer pada defer.go
	defer success()

	if error {
		panic("ada erorrrrrrrr")
	} else {
		fmt.Println("Aplikasi selesai- -----")
	}
}

func main() {
	initApp(true)
}
