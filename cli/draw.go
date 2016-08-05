package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/movement"
)

func Draw(
	frame *Frame,
) {
	// calculate the latest position
	movement.Calculate()

	width, height := termbox.Size()

	// clear
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// draw board
	DrawBoard(width, height, frame)

	// draw border
	DrawBorder(width, height, frame)

	// flush
	termbox.Flush()
}
