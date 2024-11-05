package physics

import (
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
)

type ObjectType int32

const (
	WALL = iota
	PLAYER
	ENEMY
	MINER
	COLLECTER
	MEDIC
	DEPOSIT
	KART
)

type PhysicsObject interface {
	GetType() ObjectType
	GetPosition() vector.Vector
	GetVelocity() vector.Vector
	GetCollider() *Collider
	OnCollision(other PhysicsObject)
	SetPosition(x float64, z float64)
	Draw(sdl *sdlmgr.SDLManager)
}
