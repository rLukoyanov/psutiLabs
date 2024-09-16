package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id    int
	Name  string
	Owner Person
}

func main() {
	var acc Account = Account{Id: 1, Name: "My Account", Owner: Person{Id: 1, Name: "John", Address: "123 Main St"}}

	acc.Owner = Person{1, "John", "123 Main St"}

	fmt.Println(acc.Owner.Address)
}
