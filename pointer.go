package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	// address2 := address1

	// address2.City = "Palopo"
	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Makassar", "Sulawesi Selatan", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Palopo"
	fmt.Println(address1)
	fmt.Println(address2)
}
