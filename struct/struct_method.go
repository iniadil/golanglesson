package main

import "fmt"

type User struct {
	Name string
}

// struct method sama dengan function. hanya saja function yang ada pada struct
// dan method struct hanya bisa dipanggil ketika sudah di buat/init struct nya
func (u User) sayHello() string {
	return u.Name
}

func main() {
	method := User{
		Name: "Adil Muhammad",
	}
	fmt.Println(method.sayHello())
}
