package player

import (
	"ci6450-proyecto/movement"
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	physics.PhysicsObject
	Movement *movement.Kinematic
	collider *physics.Collider
}

/*
Player constructor
*/
func New() *Player {
	var newInstance Player

	newInstance.Movement = &movement.Kinematic{
		Position: vector.Vector{
			X: 0.0,
			Z: 0.0,
		},
		Velocity: vector.Vector{
			X: 0.0,
			Z: 0.0,
		},
		Orientation: 0.0,
		Rotation:    0.0,
	}

	newInstance.collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    0.22, //Check player dimmensions later
		Height:   0.22,
	}

	return &newInstance
}

/*
Object interface implementation for player
*/

func (p *Player) GetType() physics.ObjectType {
	return physics.PLAYER
}

func (p *Player) GetPosition() vector.Vector {
	return p.Movement.Position
}

func (p *Player) GetVelocity() vector.Vector {
	return p.Movement.Velocity
}

func (p *Player) GetCollider() *physics.Collider {
	return p.collider
}

func (p *Player) OnCollision(other physics.PhysicsObject) {
	switch other.GetType() {
	case physics.WALL:
		pc := p.GetCollider()
		wc := other.GetCollider()
		const distDelta float64 = 0.015 //Small separation to prevent the make the player not touch the wall anymore

		//omg kill me
		phw := pc.Width / 2.0
		phh := pc.Height / 2.0
		whw := wc.Width / 2.0
		whh := wc.Height / 2.0

		pcx := pc.Position.X + phw
		pcz := pc.Position.Z - phh
		wcx := wc.Position.X + whw
		wcz := wc.Position.Z - whh

		diffX := pcx - wcx
		diffZ := pcz - wcz

		minXDist := phw + whw
		minZDist := phh + whh

		var depthX float64
		var depthZ float64

		if diffX > 0 {
			depthX = minXDist - diffX
		} else {
			depthX = -minXDist - diffX
		}

		if diffZ > 0 {
			depthZ = minZDist - diffZ
		} else {
			depthZ = -minZDist - diffZ
		}

		if depthX != 0 && depthZ != 0 {
			if math.Abs(depthX) < math.Abs(depthZ) {
				if depthX > 0 {
					//Left side collision
					if p.Movement.Velocity.X < 0 {
						p.Movement.Velocity.X = 0
					}
					p.Movement.Position.X = wc.Position.X + wc.Width + distDelta
				} else {
					//Right side collision
					if p.Movement.Velocity.X > 0 {
						p.Movement.Velocity.X = 0
					}
					p.Movement.Position.X = wc.Position.X - pc.Width - distDelta
				}
			} else {
				if depthZ > 0 {
					//Bottom collision
					if p.Movement.Velocity.Z < 0 {
						p.Movement.Velocity.Z = 0
					}
					p.Movement.Position.Z = wc.Position.Z + pc.Height + distDelta
				} else {
					//Top collision
					if p.Movement.Velocity.Z > 0 {
						p.Movement.Velocity.Z = 0
					}
					p.Movement.Position.Z = wc.Position.Z - wc.Height - distDelta
				}
			}
		}

	default:
	}
}

func (p *Player) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer

	playerPos := sdlmgr.FloatToPixelPos(&p.Movement.Position)
	playerSprite := sdl.Rect{
		X: playerPos.X,
		Y: playerPos.Z,
		H: 14,
		W: 14,
	}

	orientationVector := movement.OrientationAsVector(p.Movement.Orientation)
	orientationVector.Add(&p.Movement.Position)
	orientationVectorPx := sdlmgr.FloatToPixelPos(orientationVector)

	renderer.SetDrawColor(0xFF, 0xA5, 0x00, 0x00) //Player is orange
	renderer.FillRect(&playerSprite)
	renderer.SetDrawColor(0x9D, 0x00, 0xFF, 0x00) //Orientation line is purple
	renderer.DrawLine(
		playerPos.X+7,
		playerPos.Z+7,
		orientationVectorPx.X+7,
		orientationVectorPx.Z+7,
	)

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
