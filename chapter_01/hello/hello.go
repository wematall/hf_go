package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s := "Hello, Go!"     // string or series of bytes
	s1 := "head first go" // string or series of bytes
	ch := 'A'             // Rune example
	ch1 := 'b'            // Rune example

	a := math.Floor(2.75) // result of function

	fmt.Println(s)
	fmt.Println(strings.Title(s1))
	fmt.Println(a)
	fmt.Println("Rune example: ", ch)
	fmt.Println("Rune example: ", ch1)
	fmt.Println()

	fmt.Println("Hello, \nGo!")     // new line character
	fmt.Println("Hello, \tGo\t!")   // a tab character
	fmt.Println("Hello, \"a\" Go!") // double quotation mark
	fmt.Println("Hello, \\ Go!")    // a back slash
}
