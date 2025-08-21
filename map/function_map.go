package main

import "fmt"

func main() {
	// len
	// map[key]
	// map[key] = value
	// make(map[string/int/...]string
	// delete(map, "key")

	var books map[string]string = make(map[string]string)
	books["golang"] = "belajar golang"
	books["javascript"] = "java untuk orang lemah"
	books["php"] = "php untuk orang lemah"
	fmt.Println(books)
	delete(books, "golang")
	fmt.Println(books)
	fmt.Println(books["javascript"])
}
