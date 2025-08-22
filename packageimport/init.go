package main

import (
	"fmt"
	database "golang-dasar/packageimport/initialization"
	// gunakan `_` untuk melalukan init paksa
	_ "golang-dasar/packageimport/initsaja"
)

func main() {
	fmt.Println(database.GetDatabase())

}
