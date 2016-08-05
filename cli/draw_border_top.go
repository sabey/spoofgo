package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/cli/draw-termbox"
)

func DrawBorderTop(
	width int,
	height int,
	offset int,
	frame *Frame,
) {
	if _, on := frame.Options.Option_Speed.Get(); on {
		// display speed
		offset += draw.BufferHorizontal(
			width,
			height,
			offset,
			0,
			BufferSpeed(frame)...,
		)
		offset += 3 // 3 spaces after
	}
	max := width - offset
	if max >= len(buffer_menubar)+1 {
		// print all of the menubar
		draw.BufferHorizontal(
			width,
			height,
			width-len(buffer_menubar)-1,
			0,
			buffer_menubar...,
		)
	} else if max > 1 {
		// we can print atleast one character
		draw.BufferHorizontal(
			width,
			height,
			width-max,
			0,
			buffer_menubar[:max-1]..., // end with a space
		)
	}
}
