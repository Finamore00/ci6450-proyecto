package game

import (
	"ci6450-proyecto/enemy"
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/physics"
	"ci6450-proyecto/player"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

/*
Holds all game state
*/
type GameManager struct {
	Player   *player.Player
	Map      *mapa.Map
	Enemies  []*enemy.Enemy
	Physics  *physics.PhysicsManager
	Graphics *sdlmgr.SDLManager
	Running  bool
}

/*
Constructor function
*/
func New() *GameManager {
	var instance GameManager

	instance.Player = player.New()
	instance.Map = mapa.New(40, 20)
	instance.Physics = physics.NewManager()
	instance.Enemies = make([]*enemy.Enemy, 0, 25)
	instance.Graphics = sdlmgr.New()
	instance.Running = true

	return &instance
}

/*
Process general input (pause, quitting, etc.)
*/
func (g *GameManager) ProcessInput() {
	keys := g.Graphics.GetInput()

	//Quitting the game
	if keys[sdl.SCANCODE_ESCAPE] != 0 {
		g.Running = false
	}
}

/*
Update player state
*/
func (g *GameManager) UpdatePlayer(dt float64) {
	keys := g.Graphics.GetInput()
	anyInput := false

	if keys[sdl.SCANCODE_W] != 0 {
		anyInput = true
		g.Player.Movement.Velocity.Z = movement.MaxVelocity
	}
	if keys[sdl.SCANCODE_S] != 0 {
		anyInput = true
		g.Player.Movement.Velocity.Z = -movement.MaxVelocity
	}
	if keys[sdl.SCANCODE_A] != 0 {
		anyInput = true
		g.Player.Movement.Velocity.X = -movement.MaxVelocity
	}
	if keys[sdl.SCANCODE_D] != 0 {
		anyInput = true
		g.Player.Movement.Velocity.X = movement.MaxVelocity
	}

	if !anyInput {
		//If no input was detected bring player to a halt
		g.Player.Movement.Velocity.X = 0
		g.Player.Movement.Velocity.Z = 0
	}

	playerSteering := movement.SteeringOutput{
		Linear:  vector.New(0, 0),
		Angular: 0.0,
	}

	g.Player.Movement.Update(&playerSteering, dt)
	g.Player.Movement.Orientation = movement.NewOrientation(g.Player.Movement.Orientation, &g.Player.Movement.Velocity)
}

func (g *GameManager) UpdateEnemies(dt float64) {
	for _, e := range g.Enemies {
		e.EnactBehaviour(dt)
	}
}

/*
Updates all graphics on screen
*/
func (g *GameManager) UpdateGraphics() {
	g.Map.Draw(g.Graphics)
	g.Player.Draw(g.Graphics)

	for _, e := range g.Enemies {
		e.Draw(g.Graphics)
	}
}
