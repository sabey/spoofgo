package compass

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

func Coords(
	width int,
	height int,
	top_left int,
	top_right int,
	bottom_right int,
	bottom_left int,
	angle int,
) (
	int, // x
	int, // y
	uint8, // block
) {
	if angle <= top_right {
		// top right
		// get the percentage of our angle out of (top_right+(360-top_left)) degrees
		degrees := top_right + (360 - top_left)
		percent_of_angle := float64(angle) / float64(degrees)
		// get the percentage in width
		percent_in_width := int(percent_of_angle * float64(width))
		// add the percentage in width to half of width
		x := width/2 + percent_in_width
		if x == width {
			// if x equals width we have to subtract 1
			// this is an off by 1 buga
			x--
		}
		return x, // x
			0, // y = 0
			1 // block
	}
	if angle <= bottom_right {
		// right
		// subract top_right degrees from the angle
		angle -= top_right
		// get the percentage of our angle out of (bottom_right-top_right) degrees
		degrees := bottom_right - top_right
		percent_of_angle := float64(angle) / float64(degrees)
		// get the percentage in height
		percent_in_height := int(percent_of_angle * float64(height))
		if percent_in_height == height {
			// if percent_in_height equals height we have to subtract 1
			// this is an off by 1 bug
			percent_in_height--
		}
		return width - 1, // x
			percent_in_height, // y
			2 // block
	}
	if angle <= bottom_left {
		// bottom
		// subract bottom_right degrees from the angle
		angle -= bottom_right
		// get the percentage of our angle out of (bottom_left-bottom_right) degrees
		degrees := bottom_left - bottom_right
		percent_of_angle := float64(angle) / float64(degrees)
		// get the percentage in width
		percent_in_width := int(percent_of_angle * float64(width))
		x := width - percent_in_width
		if x == width {
			// if x equals width we have to subtract 1
			// this is an off by 1 bug
			x--
		}
		return x, // x
			height - 1, // y
			3 // block
	}
	if angle <= top_left {
		// left
		// subract bottom_left degrees from the angle
		angle -= bottom_left
		// get the percentage of our angle out of (top_left-bottom_left) degrees
		degrees := top_left - bottom_left
		percent_of_angle := float64(angle) / float64(degrees)
		// get the percentage in height
		percent_in_height := int(percent_of_angle * float64(height))
		y := height - percent_in_height
		if y == height {
			// if y equals height we have to add 1
			// this is an off by 1 bug
			y--
		}
		return 0, // x = 0
			y, // y
			4 // block
	}
	// top left
	angle = angle % 360 // take the modulus of 360 so we don't overflow
	// subract top_left degrees from the angle
	angle -= top_left
	// get the percentage of our angle out of (top_right+(360-top_left)) degrees
	degrees := top_right + (360 - top_left)
	// get the percentage of our angle out of 90 degrees
	percent_of_angle := float64(angle) / float64(degrees)
	// get the percentage in width
	percent_in_width := int(percent_of_angle * float64(width))
	if percent_in_width == width {
		// if percent_in_width equals width we have to subtract 1
		// this is an off by 1 bug
		percent_in_width--
	}
	return percent_in_width, // x
		0, // y = 0
		5 // block
}
