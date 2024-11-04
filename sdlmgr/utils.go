package sdlmgr

import "ci6450-proyecto/vector"

/*
Type representing a discrete pixel coordinate within the display
*/
type PixelCoordinate struct {
	X int32
	Z int32
}

/*
Turns a float position (Game coordinates) to discrete pixel positions
within the screen.
*/
func FloatToPixelPos(position *vector.Vector) PixelCoordinate {
	//Calculate line functions for coordiante-to-pixel mapping

	var result PixelCoordinate

	var horLine struct {
		m float64
		b float64
	}

	var vertLine struct {
		m float64
		b float64
	}

	horLine.m = float64(ScreenWidth-0) / (MapWidth - (-MapWidth))
	horLine.b = float64(ScreenWidth) - horLine.m*MapWidth

	vertLine.m = float64(ScreenHeight-0) / (-MapHeight - MapHeight)
	vertLine.b = float64(ScreenHeight) - vertLine.m*(-MapHeight)

	result.X = int32(horLine.m*position.X + horLine.b)
	result.Z = int32(vertLine.m*position.Z + vertLine.b)

	return result
}
