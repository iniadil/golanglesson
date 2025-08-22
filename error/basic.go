package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, operator int) (int, error) {
	if nilai == operator {
		return 0, errors.New("errors aja")
	} else {
		return nilai * nilai, nil
	}
}

func main() {
	result, err := pembagian(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
