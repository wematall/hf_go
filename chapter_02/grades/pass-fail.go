package main

import "fmt"

func main() {
	grade := 59

	if grade >= 60 {
		fmt.Println("Student grade:", grade, "This student pass exam")
	} else {
		fmt.Println("Student grade:", grade, "This student doesn't pass exam")
	}
}
