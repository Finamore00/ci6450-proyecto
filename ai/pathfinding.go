package ai

import (
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
)

type PathFinding struct {
	character  *movement.Kinematic
	target     *vector.Vector
	mapData    *mapa.Map
	Path       []vector.Vector
	currObjPos int //Position in path of the current node to go to
}

/*
Constructor for PathFinding struct
*/
func NewPathFinding(character *movement.Kinematic, mapData *mapa.Map, target *vector.Vector) *PathFinding {
	return &PathFinding{
		character:  character,
		target:     target,
		mapData:    mapData,
		Path:       nil,
		currObjPos: -1,
	}
}

/*
Clears the currently found path. Intended to be used by smart agents
when switching path-finding targets.
*/
func (p *PathFinding) ClearPath() {
	p.Path = nil
	p.currObjPos = -1
}

/*
Sets a new target for the pathFinding struct. This will cause any path generated
to the previous target to be cleared
*/
func (p *PathFinding) SetTarget(newGoal *vector.Vector) {
	if p == nil {
		return
	}

	p.target = newGoal
	p.ClearPath()
}

/*
Returns the necessary steering output to make the pathfinding
character follow the found path. The Steering Behaviour algorithm
used for following nodes is Dynamic Seek
*/
func (p *PathFinding) FollowPath() *movement.SteeringOutput {
	if p.mapData == nil {
		return nil //can't pathfind with no map dummy
	}

	if p.Path == nil {
		//If path hasn't been calculated yet. Calculate it
		p.Path = p.mapData.FindPath(p.character.Position, *p.target)
		p.currObjPos = len(p.Path) - 1
	}

	//If already at goal node do nothing
	currChNode := p.mapData.GetTileNode(p.character.Position)

	//If we're already at the current target node, switch to the next node in the path
	if currChNode == p.mapData.GetTileNode(p.Path[p.currObjPos]) && p.currObjPos > 0 {
		p.currObjPos -= 1
	}

	//Build Seek target for steering output
	stTgt := movement.Kinematic{
		Position:    p.Path[p.currObjPos],
		Velocity:    vector.Vector{X: 0, Z: 0},
		Orientation: 0,
		Rotation:    0,
	}

	return NewDynamicSeekFlee(p.character, &stTgt, true).GetSteering()

}

/*
Draw the path from the path-finding entity to the goal
*/
func (p *PathFinding) Draw(s *sdlmgr.SDLManager) {
	if p.Path == nil {
		return
	}

	renderer := s.Renderer
	renderer.SetDrawColor(0xFF, 0x00, 0x00, 0x00) //Paths are red

	//Draw first line from character to current target
	chPxPos := sdlmgr.FloatToPixelPos(&p.character.Position)
	tgtPxPos := sdlmgr.FloatToPixelPos(&p.Path[p.currObjPos])
	renderer.DrawLine(chPxPos.X, chPxPos.Z, tgtPxPos.X, tgtPxPos.Z)

	//Draw all remaining lines till goal
	for i := p.currObjPos; i > 0; i-- {
		n1PxPos := sdlmgr.FloatToPixelPos(&p.Path[i])
		n2PxPos := sdlmgr.FloatToPixelPos(&p.Path[i-1])
		renderer.DrawLine(n1PxPos.X, n1PxPos.Z, n2PxPos.X, n2PxPos.Z)
	}

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}
