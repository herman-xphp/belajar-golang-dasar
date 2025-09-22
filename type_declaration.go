package main

import "fmt"

func main() {
	type NoKTP string

	var ktpUcok NoKTP = "1234567890"

	var contoh string = "21212121212"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpUcok)
	fmt.Println(contohKtp)
}
