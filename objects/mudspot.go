package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type MudSpot struct {
	Position vector.Vector
	Width    float64
	Height   float64
}

func NewMudSpot(position vector.Vector, width float64, height float64) *MudSpot {
	return &MudSpot{
		Position: position,
		Width:    width,
		Height:   height,
	}
}

// physics.PhysicsObject interface implementation

func (m *MudSpot) GetType() physics.ObjectType {
	return physics.MUDSPOT
}

func (m *MudSpot) GetPosition() vector.Vector {
	return m.Position
}

func (m *MudSpot) GetVelocity() vector.Vector {
	return vector.Vector{X: 0.0, Z: 0.0}
}

func (m *MudSpot) GetCollider() *physics.Collider {
	return &physics.Collider{
		Position: &m.Position,
		Width:    m.Width,
		Height:   m.Height,
	}
}

func (m *MudSpot) OnCollision(other physics.PhysicsObject) {}

func (m *MudSpot) SetPosition(x float64, z float64) {
	m.Position.X = x
	m.Position.Z = z
}

func (m *MudSpot) Draw(s *sdlmgr.SDLManager) {
	r := s.Renderer

	r.SetDrawColor(0x60, 0x46, 0x0F, 0xFA) //Mud slightly transparent (?)
	pxPos := sdlmgr.FloatToPixelPos(&m.Position)
	pxDimms := sdlmgr.FloatToPixelPos(&vector.Vector{
		X: -sdlmgr.MapWidth + m.Width,
		Z: sdlmgr.MapHeight - m.Height,
	})

	r.FillRect(&sdl.Rect{
		X: pxPos.X,
		Y: pxPos.Z,
		W: pxDimms.X,
		H: pxDimms.Z,
	})

	r.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
