package main

import (
	"math"
)

type PriceFormula struct {
	MaxPrice float64
	Formula  func(float64) float64
}

var priceFormulas = []PriceFormula{
	{12000, func(price float64) float64 {
		return price + price*1.5
	}},
	{17000, func(price float64) float64 {
		return price + price*0.9
	}},
	{22000, func(price float64) float64 {
		return price + price*0.85
	}},
	{28000, func(price float64) float64 {
		return price + price*0.8
	}},
	{34000, func(price float64) float64 {
		return price + price*0.75
	}},
	{40000, func(price float64) float64 {
		return price + price*0.7
	}},
	{45000, func(price float64) float64 {
		return price + price*0.65
	}},
	{50000, func(price float64) float64 {
		return price + price*0.6
	}},
	{60000, func(price float64) float64 {
		return price + price*0.55
	}},
	{math.MaxFloat64, func(price float64) float64 {
		//для всех остальных вариантов превышающих заданные диапазоны
		return price + price*0.55
	}},
	// Добавьте другие формулы здесь
}

func SelectPriceFormula(price float64) float64 {
	for _, formula := range priceFormulas {
		if price <= formula.MaxPrice {
			//получим число по формулам
			priceAfterCalc := formula.Formula(price)
			//округлим вверх до тысяч
			priceAfterRound := math.Ceil(priceAfterCalc/1000) * 1000
			return priceAfterRound
		}
	}

	return price // Возвращаем исходную цену, если не найдена подходящая формула
}
