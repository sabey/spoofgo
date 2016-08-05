package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/compass"
	"sabey.co/spoofgo/cli/draw-termbox"
	"sabey.co/spoofgo/movement"
)

func DrawCompass(
	width int,
	height int,
	x_from int,
	y_from int,
	frame *Frame,
) {
	// black edges
	// top
	draw.PaintHorizontal(
		x_from+width,       // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		x_from,             // x_from
		x_from+width,       // x_to
		y_from,             // y_from
		' ',                // char
	)
	// left
	draw.PaintVertical(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		x_from,             // x_from
		y_from+1,           // y_from
		y_from+height-1,    // y_to
		' ',                // char
	)
	// right
	draw.PaintVertical(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		x_from+width-1,     // x_from
		y_from+1,           // y_from
		y_from+height-1,    // y_to
		' ',                // char
	)
	// bottom
	draw.PaintHorizontal(
		x_from+width,       // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		x_from,             // x_from
		x_from+width,       // x_to
		y_from+height-1,    // y_from
		' ',                // char
	)
	// draw compass
	x := 0
	y := 0
	// var block uint8
	// var colour termbox.Attribute
	if _, on := frame.Options.Option_Compass_Algorithm.Get(); on {
		// dynamic angles
		top_right, bottom_right, bottom_left, top_left := compass.CornerAngles(compass.CornerRadians(width, height))
		x, y, _ = compass.Coords(width, height, int(top_left), int(top_right), int(bottom_right), int(bottom_left), movement.GetAngle())
	} else {
		// fixed angles
		x, y, _ = compass.Coords(width, height, movement.NORTH_WEST, movement.NORTH_EAST, movement.SOUTH_EAST, movement.SOUTH_WEST, movement.GetAngle())
	}
	draw.PrintHorizontal(width, height, termbox.ColorWhite, termbox.ColorWhite, x_from+x, y_from+y, " ")
}
