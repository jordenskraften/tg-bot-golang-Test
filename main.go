package main

import (
	"fmt"
)

func main() {
	price := 150.0
	result := SelectPriceFormula(price)
	fmt.Printf("Price: %.2f, Result: %.2f\n", price, result)
}
