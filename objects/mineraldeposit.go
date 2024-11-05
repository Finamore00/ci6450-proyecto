package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

/*
Struct representing a mineral deposit. They spawn peridically at
random places in the map for miners to collect
*/
type MineralDeposit struct {
	physics.PhysicsObject
	Enabled      bool
	Location     *vector.Vector
	lastDisabled uint64
}

func NewMineralDeposit() *MineralDeposit {
	return &MineralDeposit{
		Enabled:      true,
		Location:     vector.New(0, 0),
		lastDisabled: 0,
	}
}

/*
PhysicsObject interface implementation
*/

func (m *MineralDeposit) GetType() physics.ObjectType {
	return physics.DEPOSIT
}

func (m *MineralDeposit) GetPosition() vector.Vector {
	return *m.Location
}

func (m *MineralDeposit) GetVelocity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

func (m *MineralDeposit) GetCollider() *physics.Collider {
	return &physics.Collider{
		Position: m.Location,
		Width:    0.4,
		Height:   0.4,
	}
}

func (m *MineralDeposit) SetPosition(x float64, z float64) {
	m.Location.X = x
	m.Location.Z = z
}

/*
Due to circular imports being the bane of my existence, modifications
on this object are currently handled by the other impacting party
(Miner or Collector)
*/
func (m *MineralDeposit) OnCollision(other physics.PhysicsObject) {}

func (m *MineralDeposit) Draw(s *sdlmgr.SDLManager) {
	if !m.Enabled {
		return
	}

	renderer := s.Renderer
	pixPos := sdlmgr.FloatToPixelPos(m.Location)

	//Mineral deposits are the closest I found to gold (#FFD700)
	renderer.SetDrawColor(0xFF, 0xD7, 0x00, 0xFF)
	renderer.FillRect(&sdl.Rect{
		X: pixPos.X,
		Y: pixPos.Z,
		W: 25,
		H: 25,
	})
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
