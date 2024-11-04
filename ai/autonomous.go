package ai

import (
	"ci6450-proyecto/physics"
)

type AutonomousEntity interface {
	physics.PhysicsObject
	EnactBehaviour()
}
