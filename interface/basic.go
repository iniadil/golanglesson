package main

import "fmt"

type Animal interface {
	sayHello() string
}

type AnimalType struct {
	Name string
}

// harus mengikuti tpye pada interface (nama function juga harus sama)
func (a AnimalType) sayHello() string {
	return a.Name
}

func main() {
	nama := AnimalType{
		Name: "Kucing",
	}
	fmt.Println(Animal.sayHello(nama))
}
