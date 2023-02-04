package main

import "fmt"

func main() {
	var price int = 100
	var taxRate float64 = 0.08
	var availableFunds int = 120

	var tax float64 = float64(price) * taxRate // need converion of price to float64
	var total float64 = float64(price) + tax   // need conversion of price to float64

	fmt.Println("Price is", price, "dollars.")
	fmt.Println("Tax is", tax, "dollars.")
	fmt.Println("Total cost is", total, "dollars.")
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds)) //need conversion to float64
}
