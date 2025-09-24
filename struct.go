package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", my name is", customer.Name)
}

func main() {

	var ucok Customer
	fmt.Println(ucok)

	ucok.Name = "Ucok Surucup"
	ucok.Address = "Makassar"
	ucok.Age = 24

	fmt.Println(ucok)
	fmt.Println(ucok.Name)
	fmt.Println(ucok.Address)
	fmt.Println(ucok.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     21,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Bandung", 32}
	fmt.Println(budi)

	budi.sayHello("Agus")
	ucok.sayHello("Bambang")
	joko.sayHello("Anto")
}
