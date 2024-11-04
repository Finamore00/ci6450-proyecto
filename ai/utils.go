package ai

import "math"

func MapToRange(rotation float64) float64 {
	for rotation > math.Pi {
		rotation -= 2 * math.Pi
	}
	for rotation < -math.Pi {
		rotation += 2 * math.Pi
	}
	return rotation
}
