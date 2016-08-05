package draw

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
)

type Buffer struct {
	X          int
	Y          int
	C          rune
	Foreground termbox.Attribute
	Background termbox.Attribute
}

func BufferHorizontal(
	width int,
	height int,
	x_from int,
	y_from int,
	buffer ...*Buffer,
) int {
	if y_from > height {
		// y position doesn't exist
		return 0
	}
	l := 0
	for _, b := range buffer {
		termbox.SetCell(x_from, y_from, b.C, b.Foreground, b.Background)
		x_from++
		if x_from > width {
			return l
		}
		l++
	}
	return l
}
func BufferVertical(
	width int,
	height int,
	x_from int,
	y_from int,
	buffer ...*Buffer,
) int {
	if x_from > width {
		// x position doesn't exist
		return 0
	}
	l := 0
	for _, b := range buffer {
		termbox.SetCell(x_from, y_from, b.C, b.Foreground, b.Background)
		y_from++
		if y_from > height {
			return l
		}
		l++
	}
	return l
}
