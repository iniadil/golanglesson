package main

import "fmt"

func main() {
	type NamaAsli string

	var name = "Adil"
	var OriginalName NamaAsli = NamaAsli(name) // harusnya hasilnya sama

	fmt.Println(OriginalName)
}
