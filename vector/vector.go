package vector

import "math"

/*
Data type representing a 2D vector with real (x, y) coordinates.
*/
type Vector struct {
	X float64
	Z float64
}

/*
Creates a new vector given its desired x and y values
*/
func New(x float64, z float64) *Vector {
	return &Vector{
		x,
		z,
	}
}

/*
##################################################################
######################### Struct methods #########################
##################################################################
*/

/*
Arithmetic operators. These methods modify the calling vector instance in
place
*/

func (v *Vector) Add(other *Vector) {
	v.X += other.X
	v.Z += other.Z
}

func (v *Vector) Minus(other *Vector) {
	v.X -= other.X
	v.Z -= other.Z
}

func (v *Vector) ScalarMult(n float64) {
	v.X *= n
	v.Z *= n
}

func (v *Vector) ScalarDiv(n float64) {
	v.X /= n
	v.Z /= n
}

/*
Calculates the norm of the vector v
*/
func (v *Vector) Norm() float64 {
	norm_sq := (v.X * v.X) + (v.Z * v.Z)
	return math.Sqrt(norm_sq)
}

/*
Normalizes vector v
*/
func (v *Vector) Normalize() {
	norm := v.Norm()
	if norm != 0 {
		v.ScalarDiv(norm)
	}
}

/*
Arithmetic operations. These implementations return new values
containing the operation results
*/

func Add(v1 *Vector, v2 *Vector) *Vector {
	return &Vector{
		v1.X + v2.X,
		v1.Z + v2.Z,
	}
}

func Minus(v1 *Vector, v2 *Vector) *Vector {
	return &Vector{
		v1.X - v2.X,
		v1.Z - v2.Z,
	}
}

func ScalarMult(v *Vector, n float64) *Vector {
	return &Vector{
		v.X * n,
		v.Z * n,
	}
}

func ScalarDiv(v *Vector, n float64) *Vector {
	return &Vector{
		v.X / n,
		v.Z / n,
	}
}

func DotProduct(v1 *Vector, v2 *Vector) float64 {
	return (v1.X * v2.X) + (v1.Z * v2.Z)
}
