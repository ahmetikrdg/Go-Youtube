package main

import "fmt"

func main() {
	puan := 58
	switch {
	case puan > 50 && puan < 100:
		fmt.Println("Geçti")

	case puan > 0 && puan < 50:
		fmt.Println("Kaldı")

	default:
		fmt.Println("Tekrar Deneyiniz")

	}

	i := 2

	switch i {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")
	}

}
