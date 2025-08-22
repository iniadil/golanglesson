package main

import "fmt"

type Avatar struct {
	url string
}

func main() {
	name := new(Avatar)
	name.url = "http://avatar.url"
	fmt.Println(name)
}
