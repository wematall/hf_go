package main

import (
	"fmt"
)

var originalCount int
var eatenCount int

func main() {
	originalCount = 10
	eatenCount = 4

	fmt.Println("A started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
