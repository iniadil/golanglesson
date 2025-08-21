package main

import "fmt"

func main() {
	names := "adil"

	switch names {
	case "muhammad":
		fmt.Println("muhammad")
	case "adil":
		fmt.Println("adil")
	default:
		fmt.Println("belum ada nama")
	}

	// switch short statement
	switch ages := 10; ages > 20 {
	case true:
		fmt.Println("sudah tua")
	case false:
		fmt.Println("masih muda")
	}

	// switch tanpa expression
	color := "red"
	switch {
	case len(color) > 5:
		switch color {
		case "blue":
			fmt.Println("blue")
		case "green":
			fmt.Println("green")
		default:
			fmt.Println("red")
		}
	case len(color) > 10:
		switch color {
		case "blue":
			fmt.Println("blue")
		case "green":
			fmt.Println("green")
		default:
			fmt.Println("red")
		}
	default:
		switch color {
		case "blue":
			fmt.Println("blue 1")
		case "green":
			fmt.Println("green 1")
		default:
			fmt.Println("red 1")
		}
	}
}
