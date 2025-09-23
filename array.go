package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Herman"
	names[1] = "Irna"
	names[2] = "Ucok"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		78,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))

}
