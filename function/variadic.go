package main

import "fmt"

func main() {
	fmt.Println(sumAll(50, 20, 30, 110))

	numbers := []int{50, 50, 210, 330}
	fmt.Println(sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
