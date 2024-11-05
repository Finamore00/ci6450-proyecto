package enemy

import (
	"ci6450-proyecto/ai"
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/objects"
	"ci6450-proyecto/physics"
)

type Miner struct {
	physics.PhysicsObject
	ai.AutonomousEntity
	Movement *movement.Kinematic
	Collider *physics.Collider
	mapData  *mapa.Map
	target   *objects.MineralDeposit
}
