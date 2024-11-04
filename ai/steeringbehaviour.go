package ai

import (
	"ci6450-proyecto/movement"
	"ci6450-proyecto/vector"
	"math"
)

type SteeringBehaviour interface {
	GetSteering() *movement.SteeringOutput
}

/*
Implementation of all steering behaviour classes
*/

/*
Kinematic seek and flee algorithm
*/
type KinematicSeekFlee struct {
	SteeringBehaviour
	seek      bool
	character *movement.Kinematic
	target    *movement.Kinematic
}

func (k *KinematicSeekFlee) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	//See target location and move there at full speed
	if k.seek {
		result.Linear = vector.Minus(&k.target.Position, &k.character.Position)
	} else {
		result.Linear = vector.Minus(&k.character.Position, &k.target.Position)
	}
	result.Linear.Normalize()
	result.Linear.ScalarMult(movement.MaxVelocity)
	result.Angular = 0.0

	//Make character face in velocity direction
	k.character.Orientation = movement.NewOrientation(k.character.Orientation, &k.character.Velocity)

	return &result
}

func NewKinematicSeekFlee(character *movement.Kinematic, target *movement.Kinematic, seek bool) *KinematicSeekFlee {
	return &KinematicSeekFlee{
		seek:      seek,
		character: character,
		target:    target,
	}
}

/*
Kinematic arrive algorithm
*/
type KinematicArrive struct {
	SteeringBehaviour
	character    *movement.Kinematic
	target       *movement.Kinematic
	radius       float64
	timeToTarget float64
}

func (k *KinematicArrive) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	//Get direction to target
	result.Linear = vector.Minus(&k.target.Position, &k.character.Position)

	//Check if we're within radius
	if result.Linear.Norm() < k.radius {
		return nil
	}

	//Modify velocity so arrival time matches time to target
	result.Linear.ScalarDiv(k.timeToTarget)

	//Normalize velocity result if too fast
	if result.Linear.Norm() > movement.MaxVelocity {
		result.Linear.Normalize()
		result.Linear.ScalarMult(movement.MaxVelocity)
	}

	k.character.Orientation = movement.NewOrientation(k.character.Orientation, result.Linear)

	result.Angular = 0.0
	return &result
}

func NewKinematicArrive(character *movement.Kinematic, target *movement.Kinematic) *KinematicArrive {
	return &KinematicArrive{
		character:    character,
		target:       target,
		timeToTarget: 0.25,
		radius:       1.0,
	}
}

/*
Kinematic wandering algorithm
*/
type KinematicWander struct {
	SteeringBehaviour
	character   *movement.Kinematic
	maxRotation float64
}

func (k *KinematicWander) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	orientationVector := movement.OrientationAsVector(k.character.Orientation)
	orientationVector.ScalarMult(movement.MaxVelocity)
	result.Linear = orientationVector

	result.Angular = k.maxRotation * movement.RandomBinomial()
	if angularAccel := math.Abs(result.Angular); angularAccel >= movement.MaxAngularAcceleration {
		result.Angular /= angularAccel
		result.Angular *= movement.MaxAngularAcceleration
	}

	return &result
}

func NewKinematicWander(character *movement.Kinematic) *KinematicWander {
	return &KinematicWander{
		character:   character,
		maxRotation: movement.MaxRotation,
	}
}

/*
Dynamic Seek/Flee algorithm
*/
type DynamicSeekFlee struct {
	SteeringBehaviour
	character *movement.Kinematic
	target    *movement.Kinematic
	seek      bool
}

func (d *DynamicSeekFlee) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	if d.seek {
		result.Linear = vector.Minus(&d.target.Position, &d.character.Position)
	} else {
		result.Linear = vector.Minus(&d.character.Position, &d.target.Position)
	}
	result.Linear.Normalize()
	result.Linear.ScalarMult(movement.MaxAcceleration)

	result.Angular = lookWhereYoureGoing(d.character)

	return &result
}

func NewDynamicSeekFlee(character *movement.Kinematic, target *movement.Kinematic, seek bool) *DynamicSeekFlee {
	return &DynamicSeekFlee{
		character: character,
		target:    target,
		seek:      seek,
	}
}

/*
Dynamic arrive algorithm
*/
type DynamicArrive struct {
	SteeringBehaviour
	character    *movement.Kinematic
	target       *movement.Kinematic
	targetRadius float64
	slowRadius   float64
	timeToTarget float64
}

func (d *DynamicArrive) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	direction := vector.Minus(&d.target.Position, &d.character.Position)
	distance := direction.Norm()

	if distance < d.targetRadius {
		// If we arrived at target do no steering
		return nil
	}

	var targetSpeed float64
	if distance > d.slowRadius {
		targetSpeed = movement.MaxVelocity
	} else {
		targetSpeed = movement.MaxVelocity * distance / d.slowRadius
	}

	targetVelocity := *direction
	targetVelocity.Normalize()
	targetVelocity.ScalarMult(targetSpeed)

	result.Linear = vector.Minus(&targetVelocity, &d.character.Velocity)
	result.Linear.ScalarDiv(d.timeToTarget)

	if result.Linear.Norm() > movement.MaxAcceleration {
		result.Linear.Normalize()
		result.Linear.ScalarMult(movement.MaxAcceleration)
	}

	result.Angular = lookWhereYoureGoing(d.character)
	return &result
}

func NewDynamicArriver(character *movement.Kinematic, target *movement.Kinematic) *DynamicArrive {
	return &DynamicArrive{
		character:    character,
		target:       target,
		targetRadius: 0.1,
		slowRadius:   4.0,
		timeToTarget: 0.1,
	}
}

/*
Align algorithm
*/
type Align struct {
	SteeringBehaviour
	character    *movement.Kinematic
	target       *movement.Kinematic
	targetRange  float64
	slowRange    float64
	timeToTarget float64
}

// Check why it's not working properly
func (a *Align) GetSteering() *movement.SteeringOutput {
	var result movement.SteeringOutput

	rotation := a.target.Orientation - a.character.Orientation
	rotation = MapToRange(rotation)
	rotationSize := math.Abs(rotation)

	if rotationSize < a.targetRange {
		return nil
	}

	var targetRotation float64
	if rotationSize > a.slowRange {
		targetRotation = movement.MaxRotation
	} else {
		targetRotation = movement.MaxRotation * rotationSize / a.slowRange
	}
	targetRotation *= rotation / rotationSize

	result.Angular = targetRotation - a.character.Rotation
	result.Angular /= a.timeToTarget

	if angularAccel := math.Abs(result.Angular); angularAccel > movement.MaxAngularAcceleration {
		result.Angular /= angularAccel
		result.Angular *= movement.MaxAngularAcceleration
	}

	result.Linear = vector.New(0, 0)
	return &result
}

func NewAlign(character *movement.Kinematic, target *movement.Kinematic) *Align {
	return &Align{
		character:    character,
		target:       target,
		targetRange:  math.Pi / 64,
		slowRange:    math.Pi / 8,
		timeToTarget: 0.1,
	}
}

/*
Pursue/Evade delegated behaviour algorithm
*/
type PursueEvade struct {
	SteeringBehaviour
	character     *movement.Kinematic
	target        *movement.Kinematic
	maxPrediction float64
	pursue        bool
}

func (pe *PursueEvade) GetSteering() *movement.SteeringOutput {
	var seek DynamicSeekFlee //The to-be-built target so we can delegate to seek
	seek.character = pe.character

	var direction vector.Vector
	if pe.pursue {
		direction = *vector.Minus(&pe.target.Position, &pe.character.Position)
	} else {
		direction = *vector.Minus(&pe.character.Position, &pe.target.Position)
	}
	distance := direction.Norm()

	speed := pe.character.Velocity.Norm()
	var prediction float64
	if speed <= distance/pe.maxPrediction {
		prediction = pe.maxPrediction
	} else {
		prediction = distance / speed
	}

	//Build modified target for seek
	var seekTarget movement.Kinematic
	seekTarget.Orientation = pe.target.Orientation
	seekTarget.Rotation = pe.target.Rotation
	seekTarget.Position = *vector.New(pe.target.Position.X, pe.target.Position.Z)
	seekTarget.Velocity = *vector.New(pe.target.Velocity.X, pe.target.Velocity.Z)
	seekTarget.Position.Add(vector.ScalarMult(&pe.target.Velocity, prediction))

	seek.target = &seekTarget
	seek.seek = pe.pursue

	return seek.GetSteering()
}

func NewPursueEvade(character *movement.Kinematic, target *movement.Kinematic, pursue bool) *PursueEvade {
	return &PursueEvade{
		character:     character,
		target:        target,
		pursue:        pursue,
		maxPrediction: 2.0,
	}
}

/*
Face algorithm
*/
type Face struct {
	SteeringBehaviour
	character *movement.Kinematic
	target    *movement.Kinematic
}

func (f *Face) GetSteering() *movement.SteeringOutput {
	align := NewAlign(f.character, nil) //Target will be built separately

	direction := vector.Minus(&f.target.Position, &f.character.Position)
	if direction.Norm() == 0 {
		return nil
	}

	var alignTarget movement.Kinematic
	alignTarget.Position = *vector.New(f.target.Position.X, f.target.Position.Z)
	alignTarget.Velocity = *vector.New(f.target.Velocity.X, f.target.Velocity.Z)
	alignTarget.Rotation = f.target.Rotation
	alignTarget.Orientation = math.Atan2(direction.X, direction.Z)

	align.target = &alignTarget
	return align.GetSteering()
}

func NewFace(character *movement.Kinematic, target *movement.Kinematic) *Face {
	return &Face{
		character: character,
		target:    target,
	}
}

/*
Dynamic Wander algorithm
*/
type DynamicWander struct {
	character         *movement.Kinematic
	wanderOffset      float64
	wanderRadius      float64
	wanderRate        float64
	wanderOrientation float64
}

func (dw *DynamicWander) GetSteering() *movement.SteeringOutput {
	f := NewFace(dw.character, nil) //Target will be calculated

	dw.wanderOrientation += movement.RandomBinomial() * dw.wanderRate
	targetOrientation := dw.wanderOrientation + dw.character.Orientation

	targetPos := vector.Add(&dw.character.Position, vector.ScalarMult(movement.OrientationAsVector(dw.character.Orientation), dw.wanderOffset))
	targetPos.Add(vector.ScalarMult(movement.OrientationAsVector(targetOrientation), dw.wanderRadius))

	var faceTarget movement.Kinematic
	faceTarget.Position = *targetPos
	f.target = &faceTarget

	result := f.GetSteering()
	if result == nil {
		return nil
	}

	result.Linear = vector.ScalarMult(movement.OrientationAsVector(dw.character.Orientation), movement.MaxAcceleration)

	return result
}

func NewDynamicWander(character *movement.Kinematic) *DynamicWander {
	return &DynamicWander{
		character:         character,
		wanderOffset:      1.0,
		wanderOrientation: 0.0,
		wanderRadius:      0.5,
		wanderRate:        0.2,
	}
}

/*
Look Where You're Going algorithm. Only a helper function to get orientation
since it's not really a Steering Behaviour all by itself
*/
func lookWhereYoureGoing(character *movement.Kinematic) float64 {
	velocity := character.Velocity
	if velocity.Norm() == 0 {
		return 0.0
	}

	target := movement.NewKinematic()
	target.Orientation = math.Atan2(velocity.X, velocity.Z)

	alignRes := NewAlign(character, target).GetSteering()
	if alignRes == nil {
		return 0.0
	}

	return alignRes.Angular
}
