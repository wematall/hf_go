package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	fmt.Println(input)
	fmt.Println(err)
}
