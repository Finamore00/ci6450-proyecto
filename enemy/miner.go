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
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

/*
A Miner character. It goes towards mineral deposits to mine them and
then brings the obtained ore to the mining kart for the collector to
pick up. When his stamina meter runs out he drops out cold and can't
move until the Medic tends to him.
*/
type Miner struct {
	physics.PhysicsObject
	ai.AutonomousEntity
	Movement       *movement.Kinematic
	Collider       *physics.Collider
	deposit        *objects.MineralDeposit
	kart           *objects.DepositKart
	pathFinding    *ai.PathFinding
	loaded         bool
	goingToDeposit bool //Flags
	goingToKart    bool //Flags
	stamina        float64
}

func NewMiner(mapData *mapa.Map, kart *objects.DepositKart, deposit *objects.MineralDeposit) *Miner {
	var newInstance Miner

	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    0.2,
		Height:   0.2,
	}
	newInstance.deposit = deposit
	newInstance.kart = kart
	newInstance.pathFinding = ai.NewPathFinding(newInstance.Movement, mapData, deposit.Location)
	newInstance.loaded = false
	newInstance.goingToDeposit = true
	newInstance.goingToKart = false
	newInstance.stamina = 1.0

	return &newInstance
}

/*
Implementation of PhysicsObject interface
*/

func (m *Miner) GetType() physics.ObjectType {
	return physics.MINER
}

func (m *Miner) GetPosition() vector.Vector {
	return m.Movement.Position
}

func (m *Miner) GetVelocity() vector.Vector {
	return m.Movement.Velocity
}

func (m *Miner) GetCollider() *physics.Collider {
	return m.Collider
}

func (m *Miner) SetPosition(x float64, z float64) {
	m.Movement.Position.X = x
	m.Movement.Position.Z = z
}

func (m *Miner) OnCollision(other physics.PhysicsObject) {
	switch other.GetType() {
	case physics.DEPOSIT:
		//Collision with mineral deposit
		dp, ok := other.(*objects.MineralDeposit)
		if !ok {
			fmt.Fprintln(os.Stderr, "object isn't pointer to deposit.")
			return
		}
		if dp.Enabled && m.goingToDeposit {
			m.loaded = true
			dp.Enabled = false
			dp.LastDisabled = time.Now().UnixMilli()
		}
	case physics.KART:
		kp, ok := other.(*objects.DepositKart)
		if !ok {
			fmt.Fprintln(os.Stderr, "object isn't pointer to kart.")
			return
		}

		if m.loaded {
			m.loaded = false
			kp.Load += 1
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

func (m *Miner) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer

	minerPos := sdlmgr.FloatToPixelPos(&m.Movement.Position)
	minerSprite := sdl.Rect{
		X: minerPos.X,
		Y: minerPos.Z,
		H: 14,
		W: 14,
	}

	orientationVector := movement.OrientationAsVector(m.Movement.Orientation)
	orientationVector.Add(&m.Movement.Position)
	orientationVectorPx := sdlmgr.FloatToPixelPos(orientationVector)

	renderer.SetDrawColor(0x96, 0x4B, 0x00, 0xFF) //Miner is brown
	renderer.FillRect(&minerSprite)
	renderer.SetDrawColor(0x9D, 0x00, 0xFF, 0xFF) //Orientation line is purple
	renderer.DrawLine(
		minerPos.X+7,
		minerPos.Z+7,
		orientationVectorPx.X+7,
		orientationVectorPx.Z+7,
	)

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	m.pathFinding.Draw(s) //Draw path
}

/*
ai.AutonomousEntity implementation. Consists on the miner's decision tree
*/
func (m *Miner) EnactBehaviour(dt float64) {

	if m.stamina <= 0 {
		//If exhausted, do nothing
		return
	}

	if m.loaded {
		//If carrying ore, bring it to the mining kart
		if !m.goingToKart {
			//First calculation of the path to the kart
			m.pathFinding.SetTarget(m.kart.Position)
			m.goingToKart = true
			m.goingToDeposit = false
		}
		m.Movement.Update(m.pathFinding.FollowPath(), dt)
		m.stamina -= 0.0001
	} else {
		//If deposit is active go towards it
		if m.deposit.Enabled {
			if !m.goingToDeposit {
				m.pathFinding.SetTarget(m.deposit.Location)
				m.goingToDeposit = true
				m.goingToKart = false
			}
			m.Movement.Update(m.pathFinding.FollowPath(), dt)
			m.stamina -= 0.0001
		} else {
			//If no mineral deposit is available do nothing
			return
		}
	}
}
