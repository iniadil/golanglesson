package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

func save(id string) (any, error) {
	if id == "" {
		return "", &validationError{Message: "id is empty"}
	}

	return "id nya adalah: " + id, nil
}

func main() {
	messageId, err := save("112")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(messageId)
}
