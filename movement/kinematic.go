package movement

import (
	"ci6450-proyecto/vector"
)

type Kinematic struct {
	Position    vector.Vector
	Orientation float64
	Velocity    vector.Vector
	Rotation    float64
}

/*
Update method for the Kinematic data type as described in the book
*/
func (k *Kinematic) Update(steering *SteeringOutput, time float64) {
	//If steering is nil do nothing
	if steering == nil {
		return
	}

	//Update object position and orientation
	k.Position.Add(vector.ScalarMult(&k.Velocity, time))
	k.Orientation += k.Rotation * time

	//Update object velocity and rotation
	k.Velocity.Add(vector.ScalarMult(steering.Linear, time))
	k.Rotation += steering.Angular * time

	if k.Velocity.Norm() >= MaxVelocity {
		k.Velocity.Normalize()
		k.Velocity.ScalarMult(MaxVelocity)
	}
}

/*
New Kinematic instance
*/
func NewKinematic() *Kinematic {
	return &Kinematic{
		Position:    *vector.New(0, 0),
		Velocity:    *vector.New(0, 0),
		Orientation: 0,
		Rotation:    0,
	}
}

/*
Get an equivalent Static object from the Kinematic instance
*/
func (k *Kinematic) AsStatic() *Static {
	return &Static{
		k.Position,
		k.Orientation,
	}
}
