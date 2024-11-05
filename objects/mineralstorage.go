package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type MineralStorage struct {
	physics.PhysicsObject
	Position *vector.Vector
	collider *physics.Collider
}

func NewMineralStorage() *MineralStorage {
	var newInstance MineralStorage
	newInstance.Position = vector.New(0, 0)
	newInstance.collider = &physics.Collider{
		Position: newInstance.Position,
		Width:    0.4,
		Height:   0.32,
	}

	return &newInstance
}

/*
Implementation of PhysicsObject interface
*/

func (ms *MineralStorage) GetType() physics.ObjectType {
	return physics.STORAGE
}

func (ms *MineralStorage) GetPosition() vector.Vector {
	return *ms.Position
}

func (ms *MineralStorage) GetVelocity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

func (ms *MineralStorage) SetPosition(x float64, z float64) {
	ms.Position.X = x
	ms.Position.Z = z
}

func (ms *MineralStorage) GetCollider() *physics.Collider {
	return ms.collider
}

/*
Due to circular imports being the bane of my existence, modifications
on this object are currently handled by the other impacting party
(Miner or Collector)
*/
func (ms *MineralStorage) OnCollision(other physics.PhysicsObject) {}

func (ms *MineralStorage) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer
	spritePos := sdlmgr.FloatToPixelPos(ms.Position)

	//Color of mineral storage is Burgundy (#800020)
	renderer.SetDrawColor(0x80, 0x00, 0x20, 0xFF)
	renderer.FillRect(&sdl.Rect{
		X: spritePos.X,
		Y: spritePos.Z,
		W: 25,
		H: 20,
	})
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
