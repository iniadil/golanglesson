package main

import "fmt"

// nil hampir sama dengan null pada typescript
// tapi nil tidak bisa digunakan untuk tipe data string
// nil hanya bisa untk `MAP`, `interface`, function, slice, pointer dan channel
func checkEmpty(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	name := checkEmpty("")
	if name == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("Ahlan")
	}
}
