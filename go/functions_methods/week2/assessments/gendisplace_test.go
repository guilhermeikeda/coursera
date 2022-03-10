package main

import "testing"

func TestGenDisplaceFn_Should_ReturnAFunctionToCalculateTheDisplacementBaseOnTime(testing *testing.T) {
	acceleration := float64(10)
	initialVelocity := float64(2)
	initialDisplacement := float64(1)

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	displacement := fn(3)
	if displacement != 52 {
		testing.Errorf("Invalid displacement value %f", displacement)
	}

	displacementAfterFiveSeconds := fn(5)
	if displacement != 52 {
		testing.Errorf("Invalid after 5 seconds %f", displacementAfterFiveSeconds)
	}
}
