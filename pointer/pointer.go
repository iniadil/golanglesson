package main

import "fmt"

type Customer struct {
	name string
}

func main() {
	customer := Customer{
		name: "Bob",
	}
	// dengan tanda &, maka ini adalah referensi dan bukan hasil copy paste
	customerNew := &customer
	customerNew.name = "Jane"
	fmt.Println(customer)
	fmt.Println(customerNew)

	// jika ingin membuat baru dan mengubah semuanya
	customer2 := &Customer{
		name: "nato",
	}
	customer2.name = "Sutio"
	fmt.Println(customer)
	fmt.Println(customer2)

}
