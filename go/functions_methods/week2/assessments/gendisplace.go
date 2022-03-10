package main

import (
	"fmt"
	"math"
)

func main() {
	var acceleration float64
	fmt.Println("Enter acceleration")
	fmt.Scan(&acceleration)

	var initialVelocity float64
	fmt.Println("Enter initial velocity")
	fmt.Scan(&initialVelocity)

	var initialDisplacement float64
	fmt.Println("Enter initial displacement")
	fmt.Scan(&initialDisplacement)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Println("Enter a value for time")
	fmt.Scan(&time)

	fmt.Printf("Displacement: %f", fn(time))
}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(time float64) float64 {
	displacementForTime := func(time float64) float64 {
		return (0.5 * acceleration * math.Pow(time, 2)) + (velocity * time) + displacement
	}

	return displacementForTime
}
