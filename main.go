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

	//Set walls
	g.Map.AddObstacle(vector.New(5, 2), 4, 1)

	//Set game objects
	kart := objects.NewDepositKart()
	deposit := objects.NewMineralDeposit()
	storage := objects.NewMineralStorage()
	waterSupply := objects.NewWaterSupply()

	g.Objects = append(g.Objects, kart)
	g.Objects[0].SetPosition(0.3, -4.7)
	g.Objects = append(g.Objects, deposit)
	g.Objects[1].SetPosition(7, -1)
	g.Objects = append(g.Objects, storage)
	g.Objects[2].SetPosition(-9.3, 5.0)
	g.Objects = append(g.Objects, waterSupply)
	g.Objects[3].SetPosition(0, 5.3)

	//Set game characters
	g.Enemies = append(g.Enemies, enemy.NewMiner(g.Map, kart, deposit))
	g.Enemies[0].SetPosition(-4, -3.2)
	g.Enemies = append(g.Enemies, enemy.NewCollector(g.Map, kart, storage))
	g.Enemies[1].SetPosition(-8, -3.4)
	g.Enemies = append(g.Enemies, enemy.NewMedic(g.Map, waterSupply, g.Enemies[0].(*enemy.Miner)))
	g.Enemies[2].SetPosition(-1, 5)

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
		g.Physics.CheckCollisions()
		g.Graphics.Clear()
		g.UpdateGraphics()
		g.Graphics.Render()
	}
}
