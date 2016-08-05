package draw

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
)

func PrintHorizontal(
	width int,
	height int,
	background termbox.Attribute,
	foreground termbox.Attribute,
	x_from int,
	y_from int,
	s string,
) int {
	if y_from > height {
		// y position doesn't exist
		return 0
	}
	l := 0
	for _, c := range s {
		termbox.SetCell(x_from, y_from, c, foreground, background)
		x_from++
		if x_from >= width {
			// can't print more
			return l
		}
		l++
	}
	return l
}
func PrintVertical(
	width int,
	height int,
	background termbox.Attribute,
	foreground termbox.Attribute,
	x_from int,
	y_from int,
	s string,
) int {
	if x_from > width {
		// x position doesn't exist
		return 0
	}
	l := 0
	for _, c := range s {
		termbox.SetCell(x_from, y_from, c, foreground, background)
		y_from++
		if y_from >= height {
			// can't print more
			return l
		}
		l++
	}
	return l
}
