package main

import "fmt"

func main() {
	var sehir = [2]string{"Ankara", "İstanbul"}
	fmt.Println(sehir)

	var myarr [3]string
	myarr[0] = "Antalya"
	myarr[1] = "Bursa"
	fmt.Println(myarr)

	myArray2 := [...]string{"Ali", "Veli", "Selami", "Dürdane"}
	fmt.Println(myArray2)

}
