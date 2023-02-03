package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "Hello"
	a := 42
	f := 3.14
	ch := 'A'
	bl := true

	fmt.Println(reflect.TypeOf(s))  // return string
	fmt.Println(reflect.TypeOf(a))  // return int
	fmt.Println(reflect.TypeOf(ch)) // return int32
	fmt.Println(reflect.TypeOf(bl)) // return bool
	fmt.Println(reflect.TypeOf(f))  // return float64
	fmt.Println("===========================================")

	// exercise
	// make a comment about returning type
	fmt.Println(reflect.TypeOf(25))      // return int
	fmt.Println(reflect.TypeOf(true))    // return bool
	fmt.Println(reflect.TypeOf(5.2))     // return float64
	fmt.Println(reflect.TypeOf(1))       // return int
	fmt.Println(reflect.TypeOf(false))   // return bool
	fmt.Println(reflect.TypeOf(1.0))     // return float64
	fmt.Println(reflect.TypeOf("hello")) // return string
}
