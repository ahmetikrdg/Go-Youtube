package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hey")
	}

	i := 0

	for i < 5 {
		i++
		fmt.Println("i= ", i)
	}

	var arr1 = [3]string{"Ali", "Veli", "Selami"}

	for index := range arr1 {
		fmt.Println(index)
	}

}
