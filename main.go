package main

import (
	"ci6450-proyecto/enemy"
	"ci6450-proyecto/game"
	"ci6450-proyecto/objects"
	"ci6450-proyecto/vector"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	g := game.New()
	g.Graphics.Init()

	//Set game objects
	kart := objects.NewDepositKart()
	deposit := objects.NewMineralDeposit()
	deposit.Enabled = true
	g.Map.AddObstacle(vector.New(5, 2), 4, 1)
	g.Objects = append(g.Objects, kart)
	g.Objects[0].SetPosition(0.3, -4.7)
	g.Objects = append(g.Objects, deposit)
	g.Objects[1].SetPosition(7, -1)
	g.Enemies = append(g.Enemies, enemy.NewMiner(g.Map, kart, deposit))
	g.Enemies[0].SetPosition(-3, 4)

	//Register objects in physics manager
	g.Physics.RegisterObject(g.Player)
	g.Physics.RegisterObject(g.Enemies[0])
	g.Physics.RegisterObject(g.Objects[0])
	g.Physics.RegisterObject(g.Objects[1])
	g.Map.RegisterObjects(g.Physics)

	//Initiate RNG
	rand.Seed(time.Now().UnixNano())

	//Main loop
	t := int64(0)
	for g.Running {
		for e := g.Graphics.PollEvents(); e != nil; e = g.Graphics.PollEvents() {
			if e.GetType() == sdl.QUIT {
				g.Running = false
			}
		}
		dt := time.Now().UnixMilli() - t
		t = time.Now().UnixMilli()
		g.ProcessInput()
		g.UpdatePlayer(float64(dt) / 1000)
		g.UpdateEnemies(float64(dt) / 1000)
		g.Physics.CheckCollisions()
		g.Graphics.Clear()
		g.UpdateGraphics()
		g.Graphics.Render()
	}
}
