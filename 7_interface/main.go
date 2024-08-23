package main

import "fmt"

type gasEngine struct {
	mpg     uint
	gallons uint
}

// milesLeftWithRefille implements engine.
func (g gasEngine) milesLeftWithRefille(refill uint) uint {
	return g.milesLeft() + refill
}

func (g gasEngine) milesLeft() uint {
	return g.gallons * g.mpg
}

type electricEngine struct {
	mpkwh uint
	kwh   uint
}

// milesLeftWithRefille implements engine.
func (e electricEngine) milesLeftWithRefille(refill uint) uint {
	return e.milesLeft() + refill
}

func (e electricEngine) milesLeft() uint {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint
	milesLeftWithRefille(refill uint) uint
}

func main() {
	var g1 engine = gasEngine{3, 4}
	var e1 engine = electricEngine{5, 6}
	fmt.Println(g1.milesLeft())
	fmt.Println(e1.milesLeft())
	fmt.Println(g1.milesLeftWithRefille(10))
	fmt.Println(e1.milesLeftWithRefille(10))
}
