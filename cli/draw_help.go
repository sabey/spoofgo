package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
)

func DrawHelp(
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
		"Help",
	)
	frame.Help.ResetMaybe(width-2, height-6)
	if frame.Help.HasAbove() {
		draw.PrintHorizontal(
			width,
			height,
			termbox.ColorBlack,
			termbox.ColorWhite,
			7, // x_from
			2, // y_from
			"\u2303",
		)
	}
	if frame.Help.HasBelow() {
		draw.PrintHorizontal(
			width,
			height,
			termbox.ColorBlack,
			termbox.ColorWhite,
			8, // x_from
			2, // y_from
			"\u2304",
		)
	}
	lines := frame.Help.GetContent()
	y_from := 4
	for _, line := range lines {
		draw.PrintHorizontal(
			width-2,
			height-3,
			termbox.ColorBlack,
			termbox.ColorWhite,
			2, // x_from
			y_from,
			string(line),
		)
		y_from++
	}
}
