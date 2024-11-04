package movement

//Util functions sometimes used in steering behaviours

import (
	"ci6450-proyecto/vector"
	"math"
	"math/rand"
)

/*
Returns a float value representing an orientation pointing in the
same direction as the inputted velocity
*/
func NewOrientation(current float64, velocity *vector.Vector) float64 {
	if velocity.Norm() > 0.0 {
		return math.Atan2(velocity.X, velocity.Z)
	}

	return current
}

/*
Returns a random float value between -1 and 1
*/
func RandomBinomial() float64 {
	return rand.Float64() - rand.Float64()
}

/*
Given a float orientation value returns an unit vector pointing in said
orientation
*/
func OrientationAsVector(orientation float64) *vector.Vector {
	return &vector.Vector{
		X: math.Sin(orientation),
		Z: math.Cos(orientation),
	}
}
