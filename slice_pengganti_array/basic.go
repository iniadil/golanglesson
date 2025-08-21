package main

import "fmt"

func main() {
	var array = [...]string{
		"adil",
		"muhammad",
		"zaid",
		"abdul",
		"amien",
		"huzaifi",
	}

	slice1 := array[1:3]
	fmt.Println(slice1)

	slice2 := array[:3]
	fmt.Println(slice2)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(days)

	// ketika mengubah slice data, maka array asli juga ikut berubah. kecuali append
	dayslice1 := days[4:]
	dayslice1[0] = "jumat baru"
	fmt.Println(dayslice1)
	fmt.Println(days)
}
