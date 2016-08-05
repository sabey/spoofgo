package draw

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
)

func Fill(
	width int,
	height int,
	background termbox.Attribute,
	foreground termbox.Attribute,
	x_from int,
	x_to int,
	y_from int,
	y_to int,
	c rune,
) (
	int, // x_len
	int, // y_len
) {
	if x_from > width {
		// x position doesn't exist
		return 0, 0
	}
	if x_to < x_from {
		// to was less than from
		x_to = x_from
	}
	if y_from > height {
		// y position doesn't exist
		return 0, 0
	}
	if y_to < y_from {
		// to was less than from
		y_to = y_from
	}
	for y := y_from; y < height && y < y_to; y++ {
		// move along the y axis
		for x := x_from; x < width && x < x_to; x++ {
			// move along the x axis
			termbox.SetCell(x, y, c, foreground, background)
		}
	}
	return x_to - x_from, y_to - y_from
}
