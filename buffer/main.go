package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	do_simple()
}

func do_simple() {
	sm := example_simple.Person{
		id:         12345,
		first_name: "first",
		last_name:  "last",

		is_verified: true,
		heigh:       178.5,
	}

	fmt.Println(sm)
}
