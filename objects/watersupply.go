package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type WaterSupply struct {
	physics.PhysicsObject
	Position *vector.Vector
	collider *physics.Collider
}

func NewWaterSupply() *WaterSupply {
	var newInstance WaterSupply
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

func (ws *WaterSupply) GetType() physics.ObjectType {
	return physics.WATER
}

func (ws *WaterSupply) GetPosition() vector.Vector {
	return *ws.Position
}

func (ws *WaterSupply) GetVeolcity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

func (ws *WaterSupply) SetPosition(x float64, z float64) {
	ws.Position.X = x
	ws.Position.Z = z
}

func (ws *WaterSupply) GetCollider() *physics.Collider {
	return ws.collider
}

/*
Due to circular imports being the bane of my existence, modifications
on this object are currently handled by the other impacting party
(Miner or Collector)
*/
func (ws *WaterSupply) OnCollision(other physics.PhysicsObject) {}

func (ws *WaterSupply) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer
	spritePos := sdlmgr.FloatToPixelPos(ws.Position)

	//Color of mineral storage is Sky Blue (#87CEEB)
	renderer.SetDrawColor(0x87, 0xCE, 0xEB, 0xFF)
	renderer.FillRect(&sdl.Rect{
		X: spritePos.X,
		Y: spritePos.Z,
		W: 25,
		H: 20,
	})
	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
