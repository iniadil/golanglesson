package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello := sayHelloName("Adil", spamFilter)
	fmt.Println(sayHello)
}

func sayHelloName(name string, filter func(string) string) string {
	nameFiltered := filter(name)
	return "Hallo " + nameFiltered
}

func spamFilter(name string) string {
	if strings.Contains(name, "anjing") {
		return "..."
	} else {
		return name
	}
}
