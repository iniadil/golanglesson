package main

import (
	"fmt"
	"strings"
)

type BlockedUser func(name string) bool

func main() {
	// cara pertama membuat anon func dalam variabel
	filtering := func(name string) bool {
		if strings.Contains(name, "anjing") {
			return false
		} else {
			return true
		}
	}

	fmt.Println(Greetings("Adil", filtering))

	// cara kedua langsung pada functionnya
	fmt.Println(Greetings("Adil", func(name string) bool {
		if strings.Contains(name, "anjing") {
			return false
		} else {
			return true
		}
	}))
}

func Greetings(name string, filter BlockedUser) string {
	if !filter(name) {
		return "Kamu di blokir"
	} else {
		return "Selamat Datang " + name
	}
}
