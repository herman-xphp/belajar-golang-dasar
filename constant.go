package main

import "fmt"

func main() {
	const (
		firstName string = "Ucok"
		lastName         = "Baba"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Jokowi"
}
