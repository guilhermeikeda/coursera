package main

import (
	"fmt"
	"math"
)

func mainn() {
	fmt.Println(apply(increment, 1))
	fmt.Println(apply(decrement, 1))

	var myApply func(func(int) int, int) int
	myApply = apply
	fmt.Println(myApply(increment, 10))

	anonymousFunction := func(value int) int {
		return value * 2
	}
	fmt.Println(apply(anonymousFunction, 2))

	fmt.Println(apply(func(value int) int { return value / 2 }, 4))

	printDistance()
}

func apply(function func(int) int, value int) int {
	return function(value)
}

func increment(value int) int {
	return value + 1
}

func decrement(value int) int {
	return value - 1
}

func printDistance() {
	originZero := calculateDistanceToOrigin(0, 0)
	originTwo := calculateDistanceToOrigin(2, 2)

	fmt.Println(originZero(2, 2))
	fmt.Println(originTwo(2, 2))
}

func calculateDistanceToOrigin(o_x, o_y float64) func(float64, float64) float64 {
	return func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
}
