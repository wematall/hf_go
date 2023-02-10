package main

import "fmt"

// at this example some student has a grade after exam
// if grade >= 60 student pass exam
// if grade < 60 student need take exam again

func main() {
	grade := 59

	if grade >= 60 {
		fmt.Println("Student grade:", grade, "This student pass exam")
	} else {
		fmt.Println("Student grade:", grade, "This student doesn't pass exam")
	}
}
