package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/cli/draw-termbox"
	"sabey.co/spoofgo/movement"
)

func DrawBorderLeft(
	width int,
	height int,
	offset int,
	frame *Frame,
) {
	mode, modifier := movement.GetModeModifier()
	if _, on := frame.Options.Option_Mode.Get(); on {
		_, full := frame.Options.Option_Mode_Full.Get()
		// display speed
		offset += draw.BufferVertical(
			width,
			height,
			0, // x_from
			offset,
			BufferMode(mode, full)...,
		)
		offset += 1 // space after
	}
	if _, on := frame.Options.Option_Modifier.Get(); on {
		// display modifier
		offset += draw.BufferVertical(
			width,
			height,
			0, // x_from
			offset,
			BufferModifier(modifier)...,
		)
		offset += 1 // space after
	}
}
