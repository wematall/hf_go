package main

import (
	"fmt"
	"reflect"
)

// var quantity int
// var length, width float64
// var customerName string

// another variant
var quantity = 2
var length, width = 1.2, 2.4
var customerName = "Damon Cole"

func main() {
	quantity = 2
	customerName = "Damon Cole"
	length, width = 1.2, 2.4

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	fmt.Println()

	fmt.Println(reflect.TypeOf(customerName))
	fmt.Println(reflect.TypeOf(length))
	fmt.Println(reflect.TypeOf(width))
	fmt.Println(reflect.TypeOf(quantity))
}
