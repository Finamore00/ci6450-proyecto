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

type Medic struct {
	physics.PhysicsObject
	ai.AutonomousEntity
	Movement           *movement.Kinematic
	collider           *physics.Collider
	miner              *Miner
	waterSupply        *objects.WaterSupply
	infirmaryPos       *vector.Vector
	pathFinding        *ai.PathFinding
	waterCount         uint8
	goingToMiner       bool
	goingToWaterSupply bool
	goingToInfirmary   bool
}

func NewMedic(mapData *mapa.Map, waterSupply *objects.WaterSupply, miner *Miner) *Medic {
	var newInstance Medic

	newInstance.Movement = movement.NewKinematic()
	newInstance.collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    0.2,
		Height:   0.2,
	}
	newInstance.miner = miner
	newInstance.waterSupply = waterSupply
	newInstance.infirmaryPos = vector.New(-9.0, 5.0)
	newInstance.pathFinding = ai.NewPathFinding(newInstance.Movement, mapData, newInstance.infirmaryPos)
	newInstance.waterCount = 0
	newInstance.goingToMiner = false
	newInstance.goingToWaterSupply = false
	newInstance.goingToInfirmary = true

	return &newInstance
}

/*
Implementation of PhysicsObject interface
*/

func (m *Medic) GetType() physics.ObjectType {
	return physics.MEDIC
}

func (m *Medic) GetPosition() vector.Vector {
	return m.Movement.Position
}

func (m *Medic) GetVelocity() vector.Vector {
	return m.Movement.Velocity
}

func (m *Medic) SetPosition(x float64, z float64) {
	m.Movement.Position.X = x
	m.Movement.Position.Z = z
}

func (m *Medic) GetCollider() *physics.Collider {
	return m.collider
}

func (m *Medic) OnCollision(other physics.PhysicsObject) {
	switch other.GetType() {
	case physics.WATER:
		m.waterCount = 3
	case physics.MINER:
		mp, ok := other.(*Miner)
		if !ok {
			fmt.Fprintln(os.Stderr, "object is not pointer to miner")
			return
		}
		if mp.stamina <= 0 && m.waterCount > 0 {
			mp.stamina = 1.0
			m.waterCount -= 1
		}
	case physics.WALL:
		//Wall collisions
		mc := m.GetCollider()
		wc := other.GetCollider()
		const distDelta float64 = 0.015

		mhw := mc.Width / 2.0
		mhh := mc.Height / 2.0
		whw := wc.Width / 2.0
		whh := wc.Height / 2.0

		mcx := mc.Position.X + mhw
		mcz := mc.Position.Z - mhh
		wcx := wc.Position.X + whw
		wcz := wc.Position.Z - whh

		diffX := mcx - wcx
		diffZ := mcz - wcz

		minXDist := mhw + whw
		minZDist := mhh + whh

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
					if m.Movement.Velocity.X < 0 {
						m.Movement.Velocity.X = 0
					}
					m.Movement.Position.X = wc.Position.X + wc.Width + distDelta
				} else {
					//Right side collision
					if m.Movement.Velocity.X > 0 {
						m.Movement.Velocity.X = 0
					}
					m.Movement.Position.X = wc.Position.X - mc.Width - distDelta
				}
			} else {
				if depthZ > 0 {
					//Bottom collision
					if m.Movement.Velocity.Z < 0 {
						m.Movement.Velocity.Z = 0
					}
					m.Movement.Position.Z = wc.Position.Z + mc.Height + distDelta
				} else {
					//Top collision
					if m.Movement.Velocity.Z > 0 {
						m.Movement.Velocity.Z = 0
					}
					m.Movement.Position.Z = wc.Position.Z - wc.Height - distDelta
				}
			}
		}
	default:
	}
}

func (m *Medic) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer

	medicPos := sdlmgr.FloatToPixelPos(&m.Movement.Position)
	medicSprite := sdl.Rect{
		X: medicPos.X,
		Y: medicPos.Z,
		H: 14,
		W: 14,
	}

	orientationVector := movement.OrientationAsVector(m.Movement.Orientation)
	orientationVector.Add(&m.Movement.Position)
	orientationVectorPx := sdlmgr.FloatToPixelPos(orientationVector)

	renderer.SetDrawColor(0x00, 0xFF, 0xFF, 0xFF) //Medic is cyan (#00FFFF)
	renderer.FillRect(&medicSprite)
	renderer.SetDrawColor(0x9D, 0x00, 0xFF, 0xFF) //Orientation line is purple
	renderer.DrawLine(
		medicPos.X+7,
		medicPos.Z+7,
		orientationVectorPx.X+7,
		orientationVectorPx.Z+7,
	)

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	m.pathFinding.Draw(s) //Draw path
}

/*
ai.AutonomousEntity interface implementation. Consists on the
character's behaviour tree
*/
func (m *Medic) EnactBehaviour(dt float64) {
	//If miner isn't tired. Return to the infirmary
	if m.miner.stamina > 0 {
		if !m.goingToInfirmary {
			m.pathFinding.SetTarget(m.infirmaryPos)
			m.goingToInfirmary = true
			m.goingToMiner = false
			m.goingToWaterSupply = false
		}
		//If close enough to infirmary location don't keep following
		if vector.Minus(&m.Movement.Position, m.infirmaryPos).Norm() < 0.5 {
			m.pathFinding.ClearPath()
			return
		}
		m.Movement.Update(m.pathFinding.FollowPath(), dt)
	} else {
		//If no water on stock, go to water supply
		if m.waterCount == 0 {
			if !m.goingToWaterSupply {
				m.pathFinding.SetTarget(m.waterSupply.Position)
				m.goingToWaterSupply = true
				m.goingToInfirmary = false
				m.goingToMiner = false
			}
			m.Movement.Update(m.pathFinding.FollowPath(), dt)
		} else {
			//Go tend the miner
			if !m.goingToMiner {
				m.pathFinding.SetTarget(&m.miner.Movement.Position)
				m.goingToMiner = true
				m.goingToInfirmary = false
				m.goingToWaterSupply = false
			}
			m.Movement.Update(m.pathFinding.FollowPath(), dt)
		}
	}
}
