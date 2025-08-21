package main

import "fmt"

func main() {
	nameLengkap := sayHello("Adil")
	fmt.Println(nameLengkap)
}

func sayHello(name string) string {
	return "nama kamu " + name
}
