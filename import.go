package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Ucok")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)	// tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Ucok")) // tidak bisa diakses
}
