package main

import (
	"ci6450-proyecto/enemy"
	"ci6450-proyecto/game"
	"ci6450-proyecto/objects"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	g := game.New()
	g.Player = nil
	g.Graphics.Init()

	//Set walls
	g.Map.AddObstacle(&vector.Vector{X: -sdlmgr.MapWidth, Z: sdlmgr.MapHeight}, 0.45, 2*sdlmgr.MapHeight)
	g.Map.AddObstacle(&vector.Vector{X: sdlmgr.MapWidth - 0.45, Z: sdlmgr.MapHeight}, 0.45, 2*sdlmgr.MapHeight)
	g.Map.AddObstacle(&vector.Vector{X: -sdlmgr.MapWidth, Z: sdlmgr.MapHeight}, 2*sdlmgr.MapWidth, 0.45)
	g.Map.AddObstacle(&vector.Vector{X: -sdlmgr.MapWidth, Z: -sdlmgr.MapHeight + 0.45}, 2*sdlmgr.MapWidth, 0.45)

	g.Map.AddObstacle(&vector.Vector{X: -sdlmgr.MapWidth + 0.45, Z: 2.5}, 1.2, 0.45)
	g.Map.AddObstacle(&vector.Vector{X: -6.0, Z: sdlmgr.MapHeight}, 0.45, 3.55)
	g.Map.AddObstacle(&vector.Vector{X: -7.1, Z: 2.5}, 1.2, 0.45)

	g.Map.AddObstacle(&vector.Vector{X: -2.5, Z: sdlmgr.MapHeight}, 0.45, 1.1)
	g.Map.AddObstacle(&vector.Vector{X: -6.0, Z: 2.5}, 4.0, 0.45)

	g.Map.AddObstacle(&vector.Vector{X: 7.0, Z: sdlmgr.MapHeight}, 3.0, 2.2)
	g.Map.AddObstacle(&vector.Vector{X: 8.5, Z: 3.4}, 1.5, 1.1)
	g.Map.AddObstacle(&vector.Vector{X: 6, Z: 1.1}, 4, 0.5)
	g.Map.AddObstacle(&vector.Vector{X: 4.1, Z: sdlmgr.MapHeight}, 0.55, 5.8)

	g.Map.AddObstacle(&vector.Vector{X: 3.8, Z: -1.4}, 6.8, 0.45)
	g.Map.AddObstacle(&vector.Vector{X: 3.8, Z: -1.4}, 0.6, 1.2)
	g.Map.AddObstacle(&vector.Vector{X: 8.75, Z: -4.35}, 0.9, 0.9)

	g.Map.AddObstacle(&vector.Vector{X: -5, Z: 1.1}, 2.3, 0.5)
	g.Map.AddObstacle(&vector.Vector{X: -5, Z: 1.1}, 0.5, 4.6)
	g.Map.AddObstacle(&vector.Vector{X: -5, Z: -3}, 2.3, 0.5)
	g.Map.AddObstacle(&vector.Vector{X: -4, Z: 2.2}, 0.5, 1.1)

	g.Map.AddObstacle(&vector.Vector{X: -0.5, Z: 1.1}, 2.3, 0.5)
	g.Map.AddObstacle(&vector.Vector{X: 1.3, Z: 1.1}, 0.5, 4.6)
	g.Map.AddObstacle(&vector.Vector{X: -0.5, Z: -3}, 2.3, 0.5)

	g.Map.AddObstacle(&vector.Vector{X: -7.75, Z: -0.15}, 2.75, 0.5)
	g.Map.AddObstacle(&vector.Vector{X: -sdlmgr.MapWidth, Z: -2.5}, 3.2, 0.5)

	//Set game objects
	kart := objects.NewDepositKart()
	deposit := objects.NewMineralDeposit()
	storage := objects.NewMineralStorage()
	waterSupply := objects.NewWaterSupply()

	g.Objects = append(g.Objects, kart)
	g.Objects[0].SetPosition(0.3, -4.7)
	g.Objects = append(g.Objects, deposit)
	g.Objects[1].SetPosition(5.1, 4.7)
	g.Objects = append(g.Objects, storage)
	g.Objects[2].SetPosition(-5.3, 5.0)
	g.Objects = append(g.Objects, waterSupply)
	g.Objects[3].SetPosition(-3.1, 5.0)

	//Set game characters
	g.Enemies = append(g.Enemies, enemy.NewMiner(g.Map, kart, deposit))
	g.Enemies[0].SetPosition(-8.5, 1.7)
	g.Enemies = append(g.Enemies, enemy.NewCollector(g.Map, kart, storage))
	g.Enemies[1].SetPosition(-8, -3.4)
	g.Enemies = append(g.Enemies, enemy.NewMedic(g.Map, waterSupply, g.Enemies[0].(*enemy.Miner)))
	g.Enemies[2].SetPosition(-9, 5)

	//Register objects in physics manager
	g.RegisterObjects()

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
		//Pañito caliente para el spawneo de minerales
		if !deposit.Enabled && t-deposit.LastDisabled > 30000 {
			//Choose random coordinates for the deposit
			xCoord := RandomRangeF(-sdlmgr.MapWidth, sdlmgr.MapWidth)
			zCoord := RandomRangeF(-sdlmgr.MapHeight, sdlmgr.MapHeight)
			if g.Map.QueryTiles(vector.Vector{X: xCoord, Z: zCoord}, 0.4, 0.4) {
				deposit.SetPosition(xCoord, zCoord)
				deposit.Enabled = true
			}
		}
		g.Physics.CheckCollisions()
		g.Graphics.Clear()
		g.UpdateGraphics()
		g.Graphics.Render()
	}
}

func RandomRangeF(a float64, b float64) float64 {
	return a + rand.Float64()*(b-a)
}
