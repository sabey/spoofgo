package compass

// thanks to Pablo Imeri for the corner functions

import (
	"math"
)

func CornerRadians(
	width int,
	height int,
) (
	float64, // top_right
	float64, // bottom_right
	float64, // bottom_left
	float64, // top_left
) {
	// math.Atan converts SLOPES into angles, so basically you take a slope(ie HEIGHT/WIDTH) and then it becomes an angle in radians
	top_right := math.Atan((float64(height) / 2) / (float64(width) / 2))
	// before you can find angle 2, you need to find angle 4. this is due to the nature of the tangent function.
	// it can only find angles on the RIGHT side of a circle
	top_left := (2 * math.Pi) + math.Atan((-float64(height)/2)/(float64(width)/2))
	// now that we have angle 4, we can just subtract it by a HALF rotation(which is pi, OR 180°)
	bottom_right := top_left - math.Pi
	// using the same logic we do the same thing for corner 3, but this time we're ADDING a half rotation from corner 1
	bottom_left := top_right + math.Pi
	return top_right,
		bottom_right,
		bottom_left,
		top_left
}
func CornerAngles(
	top_right float64,
	bottom_right float64,
	bottom_left float64,
	top_left float64,
) (
	float64, // top_right
	float64, // bottom_right
	float64, // bottom_left
	float64, // top_left
) {
	// corner angles are used to find the proper block
	return top_right * 180 / math.Pi,
		bottom_right * 180 / math.Pi,
		bottom_left * 180 / math.Pi,
		top_left * 180 / math.Pi
}
