package physics

import (
	"ci6450-proyecto/vector"
)

/*
Struct representing a bounding-box-based collider.
The "position" vector corresponds to the top-left corner
of the bounding box
*/
type Collider struct {
	Position *vector.Vector
	Width    float64
	Height   float64
}

/*
Function that indicates if two colliders are, in fact, colliding.
*/
func CheckCollision(c1 *Collider, c2 *Collider) bool {

	res := c1.Position.X+c1.Width >= c2.Position.X &&
		c1.Position.X <= c2.Position.X+c2.Width &&
		c1.Position.Z >= c2.Position.Z-c2.Height &&
		c1.Position.Z-c1.Height <= c2.Position.Z
	return res
}
