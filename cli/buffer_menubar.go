package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
)

var buffer_menubar = []*draw.Buffer{
	&draw.Buffer{
		C:          'E',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          's',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          'c',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          'O',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'p',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          't',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'i',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'o',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'n',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          's',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          'F',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          '8',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          'H',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'e',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'l',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'p',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          'F',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          '1',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          '2',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorRed,
	},
	&draw.Buffer{
		C:          ' ',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorBlack,
	},
	&draw.Buffer{
		C:          'N',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'e',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          'w',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
	&draw.Buffer{
		C:          's',
		Background: termbox.ColorBlack,
		Foreground: termbox.ColorWhite,
	},
}
