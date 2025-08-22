package main

import "fmt"

// interface kosong interface{} sama dengan any.
// bisa gunakan any ataupun interface{}
func sayHello() interface{} {
	//return 1
	//return "Hello"
	return 2.5
}

func main() {
	fmt.Println(sayHello())
}
