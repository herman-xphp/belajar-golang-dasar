package main

import "fmt"

func Ups() any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var empty any = Ups()
	fmt.Println(empty)
}
