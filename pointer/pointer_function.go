package main

import "fmt"

type Address struct {
	location string
}

func ChangeAddress(address *Address) {
	address.location = "Indonesia"
}

func main() {
	var address Address = Address{}
	ChangeAddress(&address)
	fmt.Println(address.location)
}
