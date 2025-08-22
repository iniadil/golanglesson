package initsaja

import "fmt"

// tanpa ada function lain, harusnya menyebabkan error karena package tidak terpakai
// gunakan `_` sebelum import untuk defends
func init() {
	fmt.Println("ini hanya akan dijalankan")
}
