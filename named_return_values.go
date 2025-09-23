package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Dhabit"
	lastName = "Masdar"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
