package main

import "fmt"

func main() {
	name := "adil"

	if name == "adil" {
		fmt.Println(true)
	} else if name == "javascript" {
		fmt.Println(false)
	} else {
		fmt.Println(name)
	}

	if name == "adil" && name == "javascript" {
		fmt.Println(true)
	}
}
