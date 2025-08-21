package main

import "fmt"

func main() {
	// konversi integer
	var nilai32 int32 = 32769
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // harusnya nilainya negatif karena maksimal value dari int16 (cek dokumentasi official)

	// konversi string
	var name string = "adil"
	var firstName = name[0] // harusnya tipe data jadi uint8
	fmt.Println(firstName)
	fmt.Println(name)
	//konversikan
	var stringConvertion = string(firstName)
	fmt.Println("Hasil konversinya = ", stringConvertion)

}
