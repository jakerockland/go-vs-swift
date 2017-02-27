package main

import (
	"fmt"
)

type CarValue interface {
	value() float64
}

type Car struct {
	Make string
	Year int
}

func (c *Car) value() float64 {
	switch {
	case c.Year >= 2000 && c.Year < 2010:
		return 6500
	case c.Year >= 2010 && c.Year < 2020:
		return 9000
	default:
		return 3000
	}
}

func NewCar(make string, year int) *Car {
	return &Car{make, year}
}

type Tesla struct {
	Car
}

func (t *Tesla) value() float64 {
	switch {
	case t.Year >= 2010 && t.Year < 2015:
		return 18000
	default:
		return 24000
	}
}

func NewTesla(year int) *Tesla {
	return &Tesla{*NewCar("Tesla", year)}
}

func value(c CarValue) float64 {
	return c.value()
}

func main() {
	herbie := NewCar("Volkswagen", 1963)
	model3 := NewTesla(2017)
	fmt.Println("Herbie is worth", value(herbie))
	fmt.Println("Model3 is worth", value(model3))
}
