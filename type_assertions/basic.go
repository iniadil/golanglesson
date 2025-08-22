package main

import "fmt"

// type assertions berguna untuk mengkonversi type any kedalam type yang di inginkan

func sayAnything() any {
	return "hello"
}

func main() {
	name := sayAnything()
	nameString := name.(string)
	fmt.Println(nameString)

	/**
	- ini tidak boleh karena akan panic atau error
	- kita harus tau dulu output nya dalam bentuk apa
	- jika tidak tahu atau ingin otomatis, gunakan expression seperti switch dengan namaFunction.(type) yang akan mengembalikan tipe data nya
	*/
	//nameInt := name.(int) // panic
	//fmt.Println(nameInt)

	// recover
	switch typeData := name.(type) {
	case int:
		fmt.Println("tipe data integer", name)
	case string:
		fmt.Println("tipe data string", name)
	default:
		fmt.Println("tipe data type", typeData)
	}

}
