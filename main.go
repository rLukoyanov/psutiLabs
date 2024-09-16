package main

import "fmt"

type f func()

func main() {
	counter := func() f {
		count := 0
		return func() {
			count += 1
			fmt.Println(count)
		}
	}

	c := counter()
	c()
	c()
	c()
	c()
	c()
	c()
	c()
	
}
