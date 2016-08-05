package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
	"sabey.co/spoofgo/movement"
)

func BufferMode(
	mode int,
	full bool,
) []*draw.Buffer {
	s := movement.PrintMode(mode, full)
	buff := make([]*draw.Buffer, 0, len(s)) // `%`
	for _, c := range s {
		buff = append(
			buff,
			&draw.Buffer{
				C:          c,
				Background: termbox.ColorBlack,
				Foreground: termbox.ColorGreen,
			},
		)
	}
	return buff
}
