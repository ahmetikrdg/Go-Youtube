package main

import "fmt"

func main() {

	m := map[int]string{
		34: "Istanbul",
		06: "Anakara",
	}

	m[35] = "İzmir"

	delete(m, 34)
	fmt.Println(m)
}
