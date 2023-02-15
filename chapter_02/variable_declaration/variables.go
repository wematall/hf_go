package main

import "fmt"

// variable declaration examples

func main() {
	// full declaration
	var a string = "hello"
	var b string
	b = "world"

	fmt.Println(a, b)

	// short declaration
	c := "Solar system"
	c, d := "Hello", "world"

	fmt.Println(c, d)
}
