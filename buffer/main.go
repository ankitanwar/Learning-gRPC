package main

import (
	right "Grpc/simple"
	"fmt"
)

func main() {
	fmt.Println("First Protocol Buffer")
	doSimple()
}

func doSimple() {
	sm := right.Person{
		Age:              22,
		FirstName:        "hello",
		LastName:         "world",
		IsProfileVerifid: true,
	}

	gg := right.Person{
		Age:              22,
		FirstName:        "hello",
		LastName:         "world",
		IsProfileVerifid: true,
	}

	fmt.Println(sm)
	fmt.Println(gg)
}
