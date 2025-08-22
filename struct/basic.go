package main

import "fmt"

type User struct {
	Name, Nip, JK string
	Age           int
}

func main() {
	var user User
	user.Name = "Muhammad Adil"
	user.Nip = "119292123"
	user.JK = "laki-laki"
	user.Age = 18

	fmt.Println(user)

	// cara kedua
	ibnu := User{
		Name: "Ibnu Khaldun",
		Nip:  "119292123",
		JK:   "laki-laki",
		Age:  18,
	}

	fmt.Println(ibnu)

	// cara ketiga
	khalisa := User{"Khalisa", "998281", "perempuan", 15}
	fmt.Println(khalisa)

	// cara akses
	fmt.Println(khalisa.Name)
	fmt.Println(khalisa.Nip)
}
