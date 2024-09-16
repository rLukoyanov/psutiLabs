package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {
	p := Person{1, "John"}

	p.UpdateName("Jane")
	fmt.Println(p.Name)

	p.SetName("Jane")
	fmt.Println(p.Name)
}
