package main

import "fmt"

func main() {
	// break akan menghentikan perulangan
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println("---- BREAK SELESAI -----")

	// continue akan menghentikan kode eksekusi dan skip ke tingkatan selanjutnya
	for a := 0; a < 10; a++ {
		if a%2 == 0 {
			fmt.Println("SKIP NOMOR ")
			continue
		}
		fmt.Println(a)
	}
	fmt.Println("---- CONTINUE SELESAI -----")
}
