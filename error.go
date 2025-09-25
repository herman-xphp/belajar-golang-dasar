package main

import (
	"errors"
	"fmt"
)

func devision(value int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Division by zero")
	} else {
		return value / divisor, nil
	}
}

func main() {
	result, err := devision(100, 0)
	if err == nil {
		fmt.Println("Result", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
