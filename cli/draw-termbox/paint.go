package draw

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
)

func PaintHorizontal(
	width int,
	height int,
	background termbox.Attribute,
	foreground termbox.Attribute,
	x_from int,
	x_to int,
	y_from int,
	c rune,
) int {
	if y_from > height {
		// y position doesn't exist
		return 0
	}
	for x := x_from; x < width && x < x_to; x++ {
		// move along the x axis
		termbox.SetCell(x, y_from, c, foreground, background)
	}
	return x_to - x_from
}
func PaintVertical(
	width int,
	height int,
	background termbox.Attribute,
	foreground termbox.Attribute,
	x_from int,
	y_from int,
	y_to int,
	c rune,
) int {
	if x_from > width {
		// x position doesn't exist
		return 0
	}
	for y := y_from; y < height && y < y_to; y++ {
		termbox.SetCell(x_from, y, c, foreground, background)
	}
	return y_to - y_from
}
