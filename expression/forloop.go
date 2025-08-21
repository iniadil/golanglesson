package main

import "fmt"

func main() {
	a := 1
	b := 1
	for a <= 10 {
		fmt.Print("*")
		for b <= 10 {
			fmt.Print("*")
			b++
		}
		fmt.Println()
		a++
	}

	// one line for expression
	for run := 1; run <= 10; run++ {
		fmt.Println("counter: ", run)
	}
	fmt.Println("------ SELESAI ------")

	// run with `slice`
	names := []string{
		"adil",
		"Muhammad",
		"Pesantren",
	}

	// method 1
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// method 2
	// range `names` adalah array/slice/map nya.
	// kalau map, maka _ adalah key nya
	for _, name := range names {
		fmt.Println(name)
	}

	// method 3
	for i, name := range names {
		fmt.Println(i, name)
	}
}
