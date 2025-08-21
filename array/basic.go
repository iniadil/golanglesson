package main

import "fmt"

func main() {
	// array golang sama seperti bahasa lain
	// ia juga memiliki index
	// ARRAY TIDAK BISA DITAMBAH DILUAR KAPASITAS YANG SUDAH DI TENTUKAN!!

	// INDEX ARRAY TIDAK BISA DIHAPUS KETIKA SUDAH DIINISIASIKAN, TETAPI BISA DI MODIFIKASI

	var names [3]string // hanya boleh berisi 3 value
	names[0] = "adil"
	names[1] = "pertama"
	fmt.Println(names) // yang ketiga akan di print kosong karena tidak diisi
	names[2] = "Muhammad"
	fmt.Println(names)
	//names[3] = "TIDAK BISA" // harusnya tidak bisa diisi karena diluar batas
}
