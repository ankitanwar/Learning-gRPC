package main

import "fmt"

func main() {
	fmt.Println("First Protocol Buffer")
	doSimple()
}

func doSimple() {
	sm := right.Person{
		age:         22,
		first_name:  "first",
		last_name:   "last",
		is_verified: true,
	}

	fmt.Println(sm)
}
