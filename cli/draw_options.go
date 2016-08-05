package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
)

func DrawOptions(
	width int,
	height int,
	frame *Frame,
) {
	draw.PrintHorizontal(
		width,
		height,
		termbox.ColorBlack,
		termbox.ColorWhite,
		2, // x_from
		2, // y_from
		"Options",
	)
	h := height - 6
	if h < 1 {
		// we want the options to still move if we can't draw them
		h = 1
	}
	frame.MenuOptions.ResetMaybe(h)
	if frame.MenuOptions.HasAbove() {
		draw.PrintHorizontal(
			width,
			height,
			termbox.ColorBlack,
			termbox.ColorWhite,
			10, // x_from
			2,  // y_from
			"\u2303",
		)
	}
	if frame.MenuOptions.HasBelow() {
		draw.PrintHorizontal(
			width,
			height,
			termbox.ColorBlack,
			termbox.ColorWhite,
			11, // x_from
			2,  // y_from
			"\u2304",
		)
	}
	item, line := frame.MenuOptions.GetItemLine()
	y_from := 4
	for ; line < len(frame.Options.Options); line++ {
		// print toggle
		c, _ := frame.Options.Options[line].Get()
		draw.PrintHorizontal(
			width-2,
			height-3,
			termbox.ColorBlack,
			termbox.ColorWhite,
			2,         // x_from
			y_from,    // y_from
			string(c), // get toggle
		)
		if line == item {
			// current option
			draw.PrintHorizontal(
				width-2,
				height-3,
				termbox.ColorBlack,
				termbox.ColorRed,
				4,      // x_from
				y_from, // y_from
				frame.Options.Options[line].GetLabel(),
			)
		} else {
			draw.PrintHorizontal(
				width-2,
				height-3,
				termbox.ColorBlack,
				termbox.ColorWhite,
				4,      // x_from
				y_from, // y_from
				frame.Options.Options[line].GetLabel(),
			)
		}
		y_from++
	}
}
