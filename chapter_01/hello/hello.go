package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "Hello, Go!"     // string or series of bytes
	s1 := "head first go" // string or series of bytes
	ch := 'A'             // Rune

	a := math.Floor(2.75) // result of function

	fmt.Println(s)
	fmt.Println(strings.Title(s1))
	fmt.Println(a)
	fmt.Println("Rune example: ", ch)
}
