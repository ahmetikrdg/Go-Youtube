package main

import "fmt"

type numbers struct {
	number1 float64
	number2 float64
}

func (n numbers) carp() float64 {
	return n.number1 * n.number2
}

type hesap interface {
	carp() float64
}

func hesapYap(h hesap) {
	fmt.Println(h.carp())
}

func main() {
	nmb := numbers{5, 5}

	hesapYap(nmb)
}
