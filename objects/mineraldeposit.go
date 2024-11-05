package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
)

/*
Struct representing a mineral deposit. They spawn peridically at
random places in the map for miners to collect
*/
type MineralDeposit struct {
	physics.PhysicsObject
	enabled  bool
	location *vector.Vector
	collider physics.Collider
}

/*
PhysicsObject interface implementation
*/

func (m *MineralDeposit) GetType() physics.ObjectType {
	return physics.DEPOSIT
}

func (m *MineralDeposit) GetPosition() vector.Vector {
	return *m.location
}

func (m *MineralDeposit) GetVelocity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

func (m *MineralDeposit) GetCollider() *physics.Collider {
	return &m.collider
}

func (m *MineralDeposit) OnCollision() {
	m.enabled = false
}

func (m *MineralDeposit) Draw(s *sdlmgr.SDLManager) {
	if !m.enabled {
		return
	}
}
