package enemy

import (
	"ci6450-proyecto/ai"
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

type EnemyType int

const (
	Steerer EnemyType = iota
	Pathfinder
)

type Enemy struct {
	physics.PhysicsObject
	ai.AutonomousEntity
	kind                  EnemyType
	Movement              *movement.Kinematic
	Collider              *physics.Collider
	steeringBehaviourData ai.SteeringBehaviour
	pathfinderData        *ai.PathFinding //Implement path finder later
}

/*
Object interface implementation for enemy
*/

func (e *Enemy) GetType() physics.ObjectType {
	return physics.ENEMY
}

func (e *Enemy) GetPosition() vector.Vector {
	return e.Movement.Position
}

func (e *Enemy) GetVelocity() vector.Vector {
	return e.Movement.Velocity
}

func (e *Enemy) GetCollider() *physics.Collider {
	return e.Collider
}

func (e *Enemy) OnCollision(other physics.PhysicsObject) {} //No collision behaviour defined yet

func (e *Enemy) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer

	enemyPos := sdlmgr.FloatToPixelPos(&e.Movement.Position)
	enemySprite := sdl.Rect{
		X: enemyPos.X,
		Y: enemyPos.Z,
		H: 14,
		W: 14,
	}

	orientationVector := movement.OrientationAsVector(e.Movement.Orientation)
	orientationVector.Add(&e.Movement.Position)
	orientationVectorPx := sdlmgr.FloatToPixelPos(orientationVector)

	renderer.SetDrawColor(0x00, 0x00, 0xFF, 0xFF) //Enemies are blue
	renderer.FillRect(&enemySprite)
	renderer.SetDrawColor(0x9D, 0x00, 0xFF, 0xFF) //Orientation line is purple
	renderer.DrawLine(
		enemyPos.X+7,
		enemyPos.Z+7,
		orientationVectorPx.X+7,
		orientationVectorPx.Z+7,
	)

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}

/*
AutonomousEntity interface implementation
*/
func (e *Enemy) EnactBehaviour(t float64) {
	switch e.kind {
	case Steerer:
		e.Movement.Update(e.steeringBehaviourData.GetSteering(), t)
	case Pathfinder:
		e.Movement.Update(e.pathfinderData.FollowPath(), t)
	default:
	}
}

/*
Setters for position and orientation
*/
func (e *Enemy) SetPosition(x float64, z float64) {
	e.Movement.Position.X = x
	e.Movement.Position.Z = z
}

func (e *Enemy) SetOrientation(angle float64) {
	e.Movement.Orientation = angle
}

/*
Constructors for each enemy type
*/
func NewKinematicSeeker(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewKinematicSeekFlee(newInstance.Movement, target, true)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewKinematicFugitive(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewKinematicSeekFlee(newInstance.Movement, target, false)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewKinematicArriver(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewKinematicArrive(newInstance.Movement, target)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewKinematicWanderer() *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewKinematicWander(newInstance.Movement)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewDynamicSeeker(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewDynamicSeekFlee(newInstance.Movement, target, true)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewDynamicFugitive(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewDynamicSeekFlee(newInstance.Movement, target, false)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewDynamicArriver(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewDynamicArriver(newInstance.Movement, target)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewAligner(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewAlign(newInstance.Movement, target)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewPursuer(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewPursueEvade(newInstance.Movement, target, true)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewEvader(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewPursueEvade(newInstance.Movement, target, false)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewFacer(target *movement.Kinematic) *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewFace(newInstance.Movement, target)
	newInstance.pathfinderData = nil

	return &newInstance
}

func NewDynamicWanderer() *Enemy {
	var newInstance Enemy

	newInstance.kind = Steerer
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    20,
		Height:   20,
	}
	newInstance.steeringBehaviourData = ai.NewDynamicWander(newInstance.Movement)
	newInstance.pathfinderData = nil

	return &newInstance
}

/*
New Path finding enemy constructor
*/
func NewPathFinder(mapData *mapa.Map, target *vector.Vector) *Enemy {
	var newInstance Enemy

	newInstance.kind = Pathfinder
	newInstance.Movement = movement.NewKinematic()
	newInstance.Collider = &physics.Collider{
		Position: &newInstance.Movement.Position,
		Width:    14,
		Height:   14, //Remember to correct these :)
	}
	newInstance.steeringBehaviourData = nil
	newInstance.pathfinderData = ai.NewPathFinding(newInstance.Movement, mapData, target)

	return &newInstance
}
