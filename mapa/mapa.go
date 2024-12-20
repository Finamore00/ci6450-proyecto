package mapa //Can't use "map" as a package name :)

import (
	"ci6450-proyecto/objects"
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"
	"math"
	"slices"

	"github.com/veandco/go-sdl2/gfx"
)

type Map struct {
	obstacles    []*objects.RegularObstacle
	mudSpots     []*objects.MudSpot
	bubbles      []*objects.OxygenBubble
	tileCols     int
	tileRows     int
	tileWidth    float64
	tileHeight   float64
	blockedTiles map[tileNode]bool
	mudTiles     map[tileNode]int
	bubbleTiles  map[tileNode]int
	Width        float64
	Height       float64 //Think later if width and height should be here
}

func New(tileColumns int, tileRows int) *Map {
	newInstance := Map{
		obstacles:    make([]*objects.RegularObstacle, 0, 10),
		mudSpots:     make([]*objects.MudSpot, 0, 10),
		bubbles:      make([]*objects.OxygenBubble, 0, 10), //It's ugly as sin but I don't care at this point man
		tileWidth:    (sdlmgr.MapWidth * 2) / float64(tileColumns),
		tileHeight:   (sdlmgr.MapHeight * 2) / float64(tileRows),
		tileCols:     tileColumns,
		tileRows:     tileRows,
		blockedTiles: make(map[tileNode]bool),
		mudTiles:     make(map[tileNode]int),
		bubbleTiles:  make(map[tileNode]int),
		Width:        sdlmgr.MapWidth,
		Height:       sdlmgr.MapHeight, //Eventually take out MapWidth and MapHeight out of sdlmgr
	}

	return &newInstance
}

func (m *Map) AddObstacle(position *vector.Vector, width float64, height float64) {
	m.obstacles = append(m.obstacles, objects.NewRegularObstacle(position, width, height))

	//Set tilegraph nodes blocked by the obstacle
	tlNode := m.GetTileNode(*position)
	trNode := m.GetTileNode(vector.Vector{X: position.X + width, Z: position.Z})
	blNode := m.GetTileNode(vector.Vector{X: position.X, Z: position.Z - height})

	for i := tlNode.X; i <= trNode.X; i++ {
		for j := tlNode.Z; j <= blNode.Z; j++ {
			m.blockedTiles[tileNode{X: i, Z: j}] = true
		}
	}

}

func (m *Map) AddMudSpot(position vector.Vector, width float64, height float64) {
	m.mudSpots = append(m.mudSpots, objects.NewMudSpot(position, width, height))

	//Set tilegraph nodes that are now muddy
	tlNode := m.GetTileNode(position)
	brNode := m.GetTileNode(vector.Vector{X: position.X + width, Z: position.Z - height})

	for i := tlNode.X; i <= brNode.X; i++ {
		for j := tlNode.Z; j <= brNode.Z; j++ {
			m.mudTiles[tileNode{X: i, Z: j}] = 1
		}
	}
}

func (m *Map) AddOxygenBubble(position vector.Vector, radius float64) {
	m.bubbles = append(m.bubbles, objects.NewOxygenBubble(position.X, position.Z, radius))

	//Set tilegraph nodes that are now oxygen rich

	//Find bubble's circumscribed square
	sideLen := radius * 2
	helperAngleTl := (3 * math.Pi) / 4
	helperAngleBr := (7 * math.Pi) / 4
	auxRad := (sideLen * math.Sqrt2) / 2

	brCorner := m.GetTileNode(vector.Vector{
		X: position.X + auxRad*math.Sin(helperAngleTl),
		Z: position.Z + auxRad*math.Cos(helperAngleTl),
	})
	tlCorner := m.GetTileNode(vector.Vector{
		X: position.X + auxRad*math.Sin(helperAngleBr),
		Z: position.Z + auxRad*math.Cos(helperAngleBr),
	})

	for i := tlCorner.X; i <= brCorner.X; i++ {
		for j := tlCorner.Z; j <= brCorner.Z; j++ {
			nodeCoord := m.GetNodeCoord(tileNode{X: i, Z: j})
			dist := vector.Minus(&position, &nodeCoord).Norm()
			if dist <= radius {
				m.bubbleTiles[tileNode{X: i, Z: j}] = 1
			}
		}
	}
}

func (m *Map) RegisterObjects(p *physics.PhysicsManager) {
	if p == nil {
		return
	}

	for _, e := range m.obstacles {
		p.RegisterObject(e)
	}

	for _, mud := range m.mudSpots {
		p.RegisterObject(mud)
	}

	for _, b := range m.bubbles {
		p.RegisterObject(b)
	}
}

func (m *Map) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer
	tileHeightHf := m.tileHeight / 2
	tileWidthHf := m.tileWidth / 2

	//Draw obstacles
	for _, o := range m.obstacles {
		o.Draw(s)
	}

	for _, mud := range m.mudSpots {
		mud.Draw(s)
	}

	for _, b := range m.bubbles {
		b.Draw(s)
	}

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)

	//Draw tiles
	renderer.SetDrawColor(0xD3, 0xD3, 0xD3, 0xFF) //Grid is light grey (#D3D3D3)
	for i := -m.Width; i <= m.Width; i += m.tileWidth {
		//Vertical lines
		topPos := sdlmgr.FloatToPixelPos(vector.New(i, m.Height))
		bottomPos := sdlmgr.FloatToPixelPos(vector.New(i, -m.Height))

		renderer.DrawLine(topPos.X, topPos.Z, bottomPos.X, bottomPos.Z)
	}

	for i := m.Height; i >= -m.Height; i -= m.tileHeight {
		//Horizontal lines
		leftPos := sdlmgr.FloatToPixelPos(vector.New(-m.Width, i))
		rightPos := sdlmgr.FloatToPixelPos(vector.New(m.Width, i))

		renderer.DrawLine(leftPos.X, leftPos.Z, rightPos.X, rightPos.Z)
	}

	//Draw representative nodes. They are apple green
	for i := -m.Width; i <= m.Width; i += m.tileWidth {
		for j := -m.Height; j <= m.Height; j += m.tileHeight {
			circleCenter := sdlmgr.FloatToPixelPos(&vector.Vector{X: i + tileWidthHf, Z: j + tileHeightHf})
			gfx.FilledCircleRGBA(renderer, circleCenter.X, circleCenter.Z, 3, 0x76, 0xCD, 0x26, 0xFF)
		}
	}

	for n := range m.mudTiles {
		nodeCoord := m.GetNodeCoord(n)
		circleCenter := sdlmgr.FloatToPixelPos(&nodeCoord)
		gfx.FilledCircleRGBA(renderer, circleCenter.X, circleCenter.Z, 4, 0xCC, 0x88, 0x99, 0xFF)

	}

	for b := range m.bubbleTiles {
		nodeCoord := m.GetNodeCoord(b)
		circleCenter := sdlmgr.FloatToPixelPos(&nodeCoord)
		gfx.FilledCircleRGBA(renderer, circleCenter.X, circleCenter.Z, 4, 0xFF, 0xBF, 0x00, 0xFF)
	}

	//Draw occupied nodes, they are red
	for k := range m.blockedTiles {
		nodeCoord := m.GetNodeCoord(k)
		circleCenter := sdlmgr.FloatToPixelPos(&nodeCoord)
		gfx.FilledCircleRGBA(renderer, circleCenter.X, circleCenter.Z, 4, 0xFF, 0x00, 0x00, 0xFF)
	}
}

/*
Path finding functionality for the map
*/

/*
Struct representing a discretized node within the tile graph. Its coordinates are integral
*/
type tileNode struct {
	X int
	Z int
}

/*
Struct representing a graph connection between two discrete graph nodes.
The costs between nodes are kept integral to avoid floating-point overhead and
inaccuracy
*/
type connection struct {
	from tileNode
	to   tileNode
	cost int
}

// Cardinal and Diagonal costs for node connections. Kept as integers for accuracy and performance
const cardCost int = 20
const diagCost int = 30

/*
Given a tileNode, returns a slice of connections containing all connections to its
neighboring nodes
*/
func (m *Map) getConnections(node tileNode, mudWeight int, bubbleWeight int) []connection {
	connections := make([]connection, 0, 8) //Any node can have at most 8 neighbors

	for h := -1; h <= 1; h++ {
		for v := -1; v <= 1; v++ {

			if h == 0 && v == 0 {
				continue //A node can't be its own neighbor
			}

			if node.X+h < 0 || node.X+h >= m.tileCols {
				continue //Node is out-of-bounds
			}

			if node.Z+v < 0 || node.Z+v >= m.tileRows {
				continue //Node is out-of-bounds
			}

			destNode := tileNode{
				X: node.X + h,
				Z: node.Z + v,
			}

			if m.blockedTiles[destNode] {
				continue //If tile is blocked don't add it
			}
			var cost int

			if h == v || h == -v {
				cost = diagCost
			} else {
				cost = cardCost
			}

			cost += mudWeight*m.mudTiles[destNode] + bubbleWeight*m.bubbleTiles[destNode]

			connections = append(connections, connection{from: node, to: destNode, cost: cost})
		}
	}
	return connections
}

/*
Given two points in the game map, finds a path between start and end using A*.
For performance reasons the returned path is reversed, so to go from start to
goal it must be followed in reverse.
*/
func (m *Map) FindPath(start vector.Vector, end vector.Vector, mudWeight int, bubbleWeight int) []vector.Vector {

	//Discretize inputted locations
	stNode := m.GetTileNode(start)
	gNode := m.GetTileNode(end)

	//Heuristic function and memoization map
	heuristicMap := map[tileNode]int{}
	octileDist := func(a tileNode) int {
		h, ok := heuristicMap[a]
		if ok {
			return h //If heursitic for this node was already calculated. Return it
		}

		dx := a.X - gNode.X
		dz := a.Z - gNode.Z

		if dx < 0 {
			dx *= -1
		}

		if dz < 0 {
			dz *= -1
		}

		if dx > dz {
			heuristicMap[a] = cardCost*dx + (diagCost-cardCost)*dz
			return heuristicMap[a]
		}

		heuristicMap[a] = cardCost*dz + (diagCost-cardCost)*dx
		return heuristicMap[a]
	}
	predecesors := map[tileNode]tileNode{} //Predecesor map. Return value

	gScore := map[tileNode]int{} //Cost-so-far for nodes
	gScore[stNode] = 0

	fScore := map[tileNode]int{} //Estimated costs to goal node
	fScore[stNode] = octileDist(stNode)

	open := make([]tileNode, 0, 100) //Open node list
	open = append(open, stNode)

	for len(open) > 0 {
		var current tileNode
		var currPos int
		minCost := math.MaxInt

		//Find the node with the lowest fScore
		for idx, n := range open {
			fs, hasFs := fScore[n]
			if hasFs && fs < minCost {
				current = n
				currPos = idx
				minCost = fs
			}
		}

		//If goal node was reached. Return predecesor map
		if current == gNode {
			//Build path with concrete map coordinates of nodes
			path := make([]vector.Vector, 0, 500)
			path = append(path, m.GetNodeCoord(gNode))
			for curr, hasPred := predecesors[gNode]; hasPred; curr, hasPred = predecesors[curr] {
				path = append(path, m.GetNodeCoord(curr))
			}
			return path
		}

		//Remove current node from open set
		open[currPos] = open[len(open)-1]
		open = open[:len(open)-1]

		for _, conn := range m.getConnections(current, mudWeight, bubbleWeight) {
			neighbor := conn.to
			gScoreCandidate := gScore[current] + conn.cost
			if currGScore, hasGScore := gScore[neighbor]; !hasGScore || gScoreCandidate < currGScore {
				predecesors[neighbor] = current
				gScore[neighbor] = gScoreCandidate
				fScore[neighbor] = gScore[neighbor] + octileDist(neighbor)

				if !slices.Contains(open, neighbor) {
					open = append(open, neighbor)
				}
			}
		}
	}

	//No path was found
	return nil
}

/*
Returns the discrete tile node corresponding to a given floating-point map coordinate.
*/
func (m *Map) GetTileNode(pos vector.Vector) tileNode {
	//Translating coordinates
	zTr := -pos.Z + m.Height
	xTr := pos.X + m.Width

	tileX := int(xTr / m.tileWidth)
	tileZ := int(zTr / m.tileHeight)

	return tileNode{
		X: tileX,
		Z: tileZ,
	}
}

/*
Returns the floating-point map coordinate of an integral graph node
*/
func (m *Map) GetNodeCoord(t tileNode) vector.Vector {
	return vector.Vector{
		X: ((float64(t.X) * m.tileWidth) - m.Width) + (m.tileWidth / 2),
		Z: (-(float64(t.Z) * m.tileHeight) + m.Height) - (m.tileHeight / 2),
	}
}

/*
Function that recieves the position of a sprite along with its dimmensions
and returns wether or not any space occupied by that object would be blocked
*/
func (m *Map) QueryTiles(position vector.Vector, width float64, height float64) bool {
	topLeft := m.GetTileNode(position)
	topRight := m.GetTileNode(*vector.New(position.X+width, position.Z))
	bottomLeft := m.GetTileNode(*vector.New(position.X, position.Z-height))

	for i := topLeft.X; i <= topRight.X; i++ {
		for j := topLeft.Z; j <= bottomLeft.Z; j++ {
			if m.blockedTiles[tileNode{X: i, Z: j}] {
				return false
			}
		}
	}

	return true
}
