package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
)

func BufferModifier(
	modifier int,
) []*draw.Buffer {
	s := fmt.Sprintf("%d", modifier)
	buff := make([]*draw.Buffer, 0, len(s)+1) // `%`
	for _, c := range s {
		buff = append(
			buff,
			&draw.Buffer{
				C:          c,
				Background: termbox.ColorBlack,
				Foreground: termbox.ColorCyan,
			},
		)
	}
	return append(
		buff,
		&draw.Buffer{
			C:          '%',
			Background: termbox.ColorBlack,
			Foreground: termbox.ColorBlue,
		},
	)
}
