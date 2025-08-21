package main

import "fmt"

func main() {
	// seperti pada typescript = {nama, lengkap} = const
	// jika salah satu value tidak dibutuhkan, ganti dengan garis _
	// ex: _, firstname := getFullName("Muhammad Adil")
	firstname, lastname := getFullName("Muhammad Adil")
	fmt.Println(firstname)
	fmt.Println(lastname)
}

func getFullName(name string) (fullName string, lastname string) {
	// ambil huruf pertama dan huruf terakhir
	return string(name[0]), string(name[len(name)-1])
}
