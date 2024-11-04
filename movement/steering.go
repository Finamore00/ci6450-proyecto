package movement

import "ci6450-proyecto/vector"

type SteeringOutput struct {
	Linear  *vector.Vector
	Angular float64
}
