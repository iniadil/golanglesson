package main

import "fmt"

type Customer struct {
	name string
}

func main() {
	customer := Customer{
		name: "Bob",
	}
	customerNew := customer
	customerNew.name = "Jane"
	fmt.Println(customer)
	fmt.Println(customerNew)

}
