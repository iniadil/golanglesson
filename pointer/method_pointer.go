package main

import "fmt"

/**
sangat dianjurkan menggunakan pointer pada method struct
*/

type Man struct {
	name string
}

// tambahkan * pada struct nya
func (m *Man) ChangeName() {
	m.name = "Mr. " + m.name
}

func main() {
	name := Man{"Adil"}
	name.ChangeName()

	fmt.Println(name.name)
}
