package main

import (
	"fmt"
	"strings"
)

// type init
type getNamaType func(name string) string

func main() {
	fmt.Println(sayHello("Adil", spamFilter))
}

func spamFilter(nama string) string {
	if strings.Contains(nama, "anjing") {
		return "..."
	} else {
		return nama
	}
}

func sayHello(nama string, filter getNamaType) string {
	return "Halo " + filter(nama)
}
