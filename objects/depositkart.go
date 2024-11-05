package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type DepositKart struct {
	physics.PhysicsObject
	Position *vector.Vector
	Load     int16
}

func NewDepositKart() *DepositKart {
	return &DepositKart{
		Position: vector.New(0, 0),
		Load:     int16(0),
	}
}

/*
Implementation of PhysicsObject interface
*/

func (d *DepositKart) GetType() physics.ObjectType {
	return physics.KART
}

func (d *DepositKart) GetPosition() vector.Vector {
	return *d.Position
}

func (d *DepositKart) GetVelocity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

func (d *DepositKart) SetPosition(x float64, z float64) {
	d.Position.X = x
	d.Position.Z = z
}

func (d *DepositKart) GetCollider() *physics.Collider {
	return &physics.Collider{
		Position: d.Position,
		Width:    0.4,
		Height:   0.32,
	}
}

/*
Due to circular imports being the bane of my existence, modifications
on this object are currently handled by the other impacting party
(Miner or Collector)
*/
func (d *DepositKart) OnCollision(other physics.PhysicsObject) {}

func (d *DepositKart) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer
	spritePos := sdlmgr.FloatToPixelPos(d.Position)

	//Color of deposit kart is Glacier Gray (#C5C6C7)
	renderer.SetDrawColor(0xC5, 0xC6, 0xC7, 0xFF)
	renderer.FillRect(&sdl.Rect{
		X: spritePos.X,
		Y: spritePos.Z,
		W: 25,
		H: 20,
	})
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
