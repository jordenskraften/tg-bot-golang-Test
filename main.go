package main

import (
	"fmt"
)

type PriceFormula struct {
	MaxPrice float64
	Formula  func(float64) float64
}

var priceFormulas = []PriceFormula{
	{135, func(price float64) float64 {
		return price + price*1.5 + 53
	}},
	{200, func(price float64) float64 {
		return price + price*0.9 + 53
	}},
	{266, func(price float64) float64 {
		return price + price*0.85 + 53
	}},
	// Добавьте другие формулы здесь
}

func selectPriceFormula(price float64) float64 {
	for _, formula := range priceFormulas {
		if price <= formula.MaxPrice {
			return formula.Formula(price)
		}
	}

	return price // Возвращаем исходную цену, если не найдена подходящая формула
}

func main() {
	price := 150.0
	result := selectPriceFormula(price)
	fmt.Printf("Price: %.2f, Result: %.2f\n", price, result)
}
