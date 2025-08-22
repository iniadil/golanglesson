package helper

/*
*
- penggunakan public function harus diawali dengan HURUF BESAR agar bisa diakses oleh package lain
- jika menggunakan huruf kecil, maka hanya package tersebut yang bisa menggunakannya
- termasuk variabel
*/

var applicationName = "Golang Dasar" // tidak bisa di akses diluar package
var Version = "1.0.0"                // bisa di akses di luar package

// bisa diakses public
func Greetings(name string) string {
	return "Hello " + name
}
