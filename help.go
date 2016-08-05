package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/textblock"
)

const (
	help = `Movement:

Accelerate:	W, Arrow Up
Left:		A, Arrow Left
Decelerate:	S, Arrow Down
Right:		D, Arrow Right

Angles:
NorthWest:	Q
NorthEast:	E
SouthWest:	Z
Flip:		X
SouthEast:	C

Mode:
Toggle:		~
Set:		1-4

Mode Modifier:
Increase:	R, =, +, Page Up
Decrease:	F, -, _, Page Down
Reset:		V, 0`
)

var (
	textblock_help *textblock.TextBlock
)

func init() {
	textblock_help = textblock.Create([]byte(help))
}
