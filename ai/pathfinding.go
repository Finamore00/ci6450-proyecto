package ai

import (
	"ci6450-proyecto/mapa"
	"ci6450-proyecto/movement"
	"ci6450-proyecto/vector"
)

type PathFinding struct {
	character  *movement.Kinematic
	target     *vector.Vector
	mapData    *mapa.Map
	path       []vector.Vector
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
		path:       nil,
		currObjPos: -1,
	}
}

/*
Clears the currently found path. Intended to be used by smart agents
when switching path-finding targets.
*/
func (p *PathFinding) ClearPath() {
	p.path = nil
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

	if p.path == nil {
		//If path hasn't been calculated yet. Calculate it
		p.path = p.mapData.FindPath(p.character.Position, *p.target)
		p.currObjPos = len(p.path) - 1
	}

	//If already at goal node do nothing
	currChNode := p.mapData.GetTileNode(p.character.Position)
	gNode := p.mapData.GetTileNode(*p.target)
	if currChNode == gNode {
		return nil
	}

	//If we're already at the current target node, switch to the next node in the path
	if currChNode == p.mapData.GetTileNode(p.path[p.currObjPos]) {
		p.currObjPos -= 1
	}

	//Build Seek target for steering output
	skTgt := movement.Kinematic{
		Position:    p.path[p.currObjPos],
		Velocity:    vector.Vector{X: 0, Z: 0},
		Orientation: 0,
		Rotation:    0,
	}

	return NewDynamicSeekFlee(p.character, &skTgt, true).GetSteering()

}
