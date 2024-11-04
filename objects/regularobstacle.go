package objects

import (
	"ci6450-proyecto/physics"
	"ci6450-proyecto/sdlmgr"
	"ci6450-proyecto/vector"

	"github.com/veandco/go-sdl2/sdl"
)

/*
Regular obstacles are rectangular structures that
don't move, but have collision logic. Basically walls
and pillars within the map
*/
type RegularObstacle struct {
	physics.PhysicsObject
	position *vector.Vector
	width    float64
	height   float64
}

func NewRegularObstacle(position *vector.Vector, width float64, height float64) *RegularObstacle {
	return &RegularObstacle{
		position: position,
		width:    width,
		height:   height,
	}
}

/*
Object interface implementation for regular obstacle struct
*/

func (o *RegularObstacle) GetType() physics.ObjectType {
	return physics.WALL
}

func (o *RegularObstacle) GetPosition() vector.Vector {
	return *o.position
}

func (o *RegularObstacle) GetVelocity() vector.Vector {
	return vector.Vector{
		X: 0.0,
		Z: 0.0,
	}
}

func (o *RegularObstacle) GetCollider() *physics.Collider {
	return &physics.Collider{
		Position: o.position,
		Width:    o.width,
		Height:   o.height,
	}
}

func (o *RegularObstacle) Draw(s *sdlmgr.SDLManager) {
	renderer := s.Renderer
	renderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF) //Walls are white

	ulPos := sdlmgr.FloatToPixelPos(o.position) //Upper-left corner position
	pxLen := sdlmgr.FloatToPixelPos(vector.New(-sdlmgr.MapWidth+o.width, sdlmgr.MapHeight))
	pxHei := sdlmgr.FloatToPixelPos(vector.New(-sdlmgr.MapWidth, sdlmgr.MapHeight-o.height))

	renderer.FillRect(&sdl.Rect{
		X: ulPos.X,
		Y: ulPos.Z,
		W: pxLen.X,
		H: pxHei.Z,
	})

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0x00)
}

func (o *RegularObstacle) OnCollision(other physics.PhysicsObject) {} //Collisions don't modify regular obstacles
