package enemy

import (
	"ci6450-proyecto/ai"
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/objects"
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
	"fmt"
	"math"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Collector struct {
	physics.PhysicsObject
	ai.AutonomousEntity
	Movement       *movement.Kinematic
	collider       *physics.Collider
	storage        *objects.MineralStorage
	kart           *objects.DepositKart
	pathFinding    *ai.PathFinding
	loaded         bool
	goingToKart    bool
	goingToStorage bool
}

func NewCollector(mapData *mapa.Map, kart *objects.DepositKart, storage *objects.MineralStorage) *Collector {
	var newInstance Collector

	newInstance.Movement = movement.NewKinematic()
	newInstance.collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    0.2,
		Height:   0.2,
	}
	newInstance.kart = kart
	newInstance.storage = storage
	newInstance.pathFinding = ai.NewPathFinding(newInstance.Movement, mapData, kart.Position)
	newInstance.loaded = false
	newInstance.goingToKart = true
	newInstance.goingToStorage = false

	return &newInstance
}

/*
Implementation of PhysicsObject interface
*/
func (c *Collector) GetType() physics.ObjectType {
	return physics.COLLECTOR
}

func (c *Collector) GetPosition() vector.Vector {
	return c.Movement.Position
}

func (c *Collector) GetVelocity() vector.Vector {
	return c.Movement.Velocity
}

func (c *Collector) SetPosition(x float64, z float64) {
	c.Movement.Position.X = x
	c.Movement.Position.Z = z
}

func (c *Collector) GetCollider() *physics.Collider {
	return c.collider
}

func (c *Collector) OnCollision(other physics.PhysicsObject) {
	switch other.GetType() {
	case physics.KART:
		//Collision with kart
		kp, ok := other.(*objects.DepositKart)
		if !ok {
			fmt.Fprintln(os.Stderr, "other object is not pointer to kart")
		}
		if kp.Load > 0 {
			c.loaded = true
			kp.Load -= 1
		}
	case physics.STORAGE:
		if c.loaded {
			c.loaded = false
		}
	case physics.WALL:
		//Wall collisions. Eventually put this into its own function
		cc := c.GetCollider()
		wc := other.GetCollider()
		const distDelta float64 = 0.015

		chw := cc.Width / 2.0
		chh := cc.Height / 2.0
		whw := wc.Width / 2.0
		whh := wc.Height / 2.0

		ccx := cc.Position.X + chw
		ccz := cc.Position.Z - chh
		wcx := wc.Position.X + whw
		wcz := wc.Position.Z - whh

		diffX := ccx - wcx
		diffZ := ccz - wcz

		minXDist := chw + whw
		minZDist := chh + whh

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
					if c.Movement.Velocity.X < 0 {
						c.Movement.Velocity.X = 0
					}
					c.Movement.Position.X = wc.Position.X + wc.Width + distDelta
				} else {
					//Right side collision
					if c.Movement.Velocity.X > 0 {
						c.Movement.Velocity.X = 0
					}
					c.Movement.Position.X = wc.Position.X - cc.Width - distDelta
				}
			} else {
				if depthZ > 0 {
					//Bottom collision
					if c.Movement.Velocity.Z < 0 {
						c.Movement.Velocity.Z = 0
					}
					c.Movement.Position.Z = wc.Position.Z + cc.Height + distDelta
				} else {
					//Top collision
					if c.Movement.Velocity.Z > 0 {
						c.Movement.Velocity.Z = 0
					}
					c.Movement.Position.Z = wc.Position.Z - wc.Height - distDelta
				}
			}
		}
	default:
	}
}

func (c *Collector) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer

	minerPos := sdlmgr.FloatToPixelPos(&c.Movement.Position)
	minerSprite := sdl.Rect{
		X: minerPos.X,
		Y: minerPos.Z,
		H: 14,
		W: 14,
	}

	orientationVector := movement.OrientationAsVector(c.Movement.Orientation)
	orientationVector.Add(&c.Movement.Position)
	orientationVectorPx := sdlmgr.FloatToPixelPos(orientationVector)

	renderer.SetDrawColor(0xF2, 0x85, 0x00, 0xFF) //Collector is tangerine (#F28500)
	renderer.FillRect(&minerSprite)
	renderer.SetDrawColor(0x9D, 0x00, 0xFF, 0xFF) //Orientation line is purple
	renderer.DrawLine(
		minerPos.X+7,
		minerPos.Z+7,
		orientationVectorPx.X+7,
		orientationVectorPx.Z+7,
	)

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	c.pathFinding.Draw(s) //Draw path
}

/*
ai.AutonomousEntity interface implementatio. Corresponds to the characters decision tree
*/
func (c *Collector) EnactBehaviour(dt float64) {
	if c.loaded {
		//If carrying a load. Take it to storage
		if !c.goingToStorage {
			c.pathFinding.SetTarget(c.storage.Position)
			c.goingToStorage = true
			c.goingToKart = false
		}
		c.Movement.Update(c.pathFinding.FollowPath(), dt)
	} else {
		//If kart has ore, go get it
		if c.kart.Load > 0 {
			if !c.goingToKart {
				c.pathFinding.SetTarget(c.kart.Position)
				c.goingToKart = true
				c.goingToStorage = false
			}
			c.Movement.Update(c.pathFinding.FollowPath(), dt)
		} else {
			//Wander around
			c.pathFinding.ClearPath()
			c.Movement.Update(ai.NewDynamicWander(c.Movement).GetSteering(), dt)
		}
	}
}
