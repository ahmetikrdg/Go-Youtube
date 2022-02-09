package main

import "fmt"

type Person struct {
	name      string
	age       int
	isMarried bool
}

func main() {

	var p1 Person
	p1.name = "Ali"
	p1.age = 25
	p1.isMarried = false

	fmt.Println(p1.name, " ", p1.age, " ", p1.isMarried)

	fmt.Println(Person{"Mehmet", 26, true})

}
