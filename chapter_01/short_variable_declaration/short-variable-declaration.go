package main

import "fmt"

func main() {
	customerName := "Damon Cole"
	quantity := 4
	length, width := 1.2, 2.4

	// we can't declare variable the same name twice
	// quantity := 4
	// we can't assign another type to variables
	// quantity = "a"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

}
