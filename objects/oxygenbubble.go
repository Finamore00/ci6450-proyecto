package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
	"math"

	"github.com/veandco/go-sdl2/gfx"
)

type OxygenBubble struct {
	physics.PhysicsObject
	Position vector.Vector
	Radius   float64
}

func NewOxygenBubble(x float64, z float64, r float64) *OxygenBubble {
	return &OxygenBubble{
		Position: vector.Vector{X: x, Z: z},
		Radius:   r,
	}
}

// physics.PhysicsObject interface implementation

func (o *OxygenBubble) GetType() physics.ObjectType {
	return physics.OXYGENBUBBLE
}

func (o *OxygenBubble) GetPosition() vector.Vector {
	return o.Position
}

func (o *OxygenBubble) GetVelocity() vector.Vector {
	return vector.Vector{X: 0, Z: 0}
}

/*
Since the physics engine doesn't support sphere colliders yet, the collider
for the spherical oxygen bubble will be represented by its inscribed square
shape.
*/
func (o *OxygenBubble) GetCollider() *physics.Collider {
	auxAngle := (3 * math.Pi) / 4.0
	sideLen := o.Radius * math.Sqrt2

	collPos := vector.Vector{
		X: o.Position.X + o.Radius*math.Sin(auxAngle),
		Z: o.Position.Z + o.Radius*math.Cos(auxAngle),
	}

	return &physics.Collider{
		Position: &collPos,
		Width:    sideLen,
		Height:   sideLen,
	}
}

func (o *OxygenBubble) OnCollision(other physics.PhysicsObject) {}

func (o *OxygenBubble) SetPosition(x float64, z float64) {
	o.Position.X = x
	o.Position.Z = z
}

func (o *OxygenBubble) Draw(s *sdlmgr.SDLManager) {
	r := s.Renderer

	circleCenter := sdlmgr.FloatToPixelPos(&o.Position)
	pxRad := sdlmgr.FloatToPixelPos(&vector.Vector{
		X: -sdlmgr.MapWidth + o.Radius,
		Z: 0.0,
	}).X

	//Oxygen bubbles are sightly transparent
	gfx.FilledCircleRGBA(
		r,
		circleCenter.X,
		circleCenter.Z,
		pxRad,
		0xC9,
		0xE8,
		0xF2,
		0xAF,
	)

	r.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
