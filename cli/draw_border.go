package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
)

const (
	NAME_HORIZONTAL = `Spoofgo.com`
	NAME_VERTICAL   = `poof.go`
)

func DrawBorder(
	width int,
	height int,
	frame *Frame,
) {
	// black edges
	// top
	draw.PaintHorizontal(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		0,                  // x_from
		width,              // x_to
		0,                  // y_from
		' ',                // char
	)
	// left
	draw.PaintVertical(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		0,                  // x_from
		1,                  // y_from
		height-1,           // y_to
		' ',                // char
	)
	// right
	draw.PaintVertical(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		width-1,            // x_from
		1,                  // y_from
		height-1,           // y_to
		' ',                // char
	)
	// bottom
	draw.PaintHorizontal(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorBlack, // foreground
		0,                  // x_from
		width,              // x_to
		height-1,           // y_from
		' ',                // char
	)

	// print name horizontal
	draw.PrintHorizontal(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorWhite, // foreground
		0,                  // x_from
		0,                  // y_from
		NAME_HORIZONTAL,    // text
	)
	top_offset := len(NAME_HORIZONTAL) + 1 // `.`
	// print project vertical
	draw.PrintVertical(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorWhite, // foreground
		0,                  // x_from
		1,                  // y_from
		NAME_VERTICAL,      // text
	)
	left_offset := len(NAME_VERTICAL) + 2 // `S.`
	// version
	draw.PrintHorizontal(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorWhite, // foreground
		top_offset,         // x_from
		0,                  // y_from
		"(v",               // text
	)
	top_offset += 2 // `(v`
	version := GetVersion(frame)
	draw.PrintHorizontal(
		width,                   // width
		height,                  // height
		termbox.ColorBlack,      // background
		GetVersionColour(frame), // foreground
		top_offset,              // x_from
		0,                       // y_from
		version,                 // text
	)
	top_offset += len(version)
	draw.PrintHorizontal(
		width,              // width
		height,             // height
		termbox.ColorBlack, // background
		termbox.ColorWhite, // foreground
		top_offset,         // x_from
		0,                  // y_from
		")",                // text
	)
	top_offset += 2 // `).`
	// left
	DrawBorderLeft(width, height, left_offset, frame)
	// top
	DrawBorderTop(width, height, top_offset, frame)
}
