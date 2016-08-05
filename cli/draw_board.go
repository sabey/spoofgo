package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"sabey.co/spoofgo/cli/draw-termbox"
	"sabey.co/spoofgo/duration"
	"sabey.co/spoofgo/movement"
	"time"
)

func DrawBoard(
	width int,
	height int,
	frame *Frame,
) {
	x_from := 1
	y_from := 1
	// draw canvas
	fill_width, fill_height := draw.Fill(
		width,                // width
		height,               // height
		termbox.ColorDefault, // background
		termbox.ColorBlack,   // foreground
		x_from,               // x_from
		width-1,              // x_to
		y_from,               // y_from
		height-1,             // y_to,
		' ',                  // char
	)
	// DRAW LOCATION
	if frame.Location == OPTIONS ||
		frame.Location == HELP ||
		frame.Location == NEWS {
		// draw alert
		if _, on := frame.Options.Option_Alert.Get(); on {
			// alert the user they're outside of player mode
			// top/bottom
			for x := 1; x < width-1; x++ {
				termbox.SetCell(x, 1, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
				termbox.SetCell(x, height-2, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
			}
			// left/right
			for y := 2; y < height-2; y++ {
				termbox.SetCell(1, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
				termbox.SetCell(width-2, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
			}
		}
	}
	if frame.Location == OPTIONS {
		// DRAW OPTIONS
		DrawOptions(width, height, frame)
		return
	}
	if frame.Location == HELP {
		// DRAW HELP
		DrawHelp(width, height, frame)
		return
	}
	if frame.Location == NEWS {
		// DRAW NEWS
		DrawNews(width, height, frame)
		return
	}
	// DRAW BOARD
	if _, on := frame.Options.Option_Compass.Get(); on {
		// Compass
		DrawCompass(fill_width, fill_height, x_from, y_from, frame)
		x_from++
		y_from++
		fill_width -= 2
		fill_height -= 2
	}
	s := movement.GetState()
	if _, on := frame.Options.Option_Center_Position.Get(); on {
		// Center
		colour := termbox.ColorCyan
		if _, on := frame.Options.Option_Center_Position_Moving.Get(); on {
			if s.Accelerating {
				colour = termbox.ColorGreen
			} else {
				colour = termbox.ColorRed
			}
		}
		draw.PrintHorizontal(
			fill_width,             // width,
			fill_height,            // height,
			colour,                 // background,
			termbox.ColorWhite,     // foreground,
			x_from+(fill_width/2),  // x_from,
			y_from+(fill_height/2), // y_from,
			" ", // s,
		)
	}
	if _, on := frame.Options.Option_Compass_North.Get(); on {
		// North
		draw.PrintHorizontal(
			fill_width,            // width,
			fill_height,           // height,
			termbox.ColorBlack,    // background,
			termbox.ColorWhite,    // foreground,
			x_from+(fill_width/2), // x_from,
			y_from, // y_from,
			"N",    // s,
		)
	}
	if _, on := frame.Options.Option_Angle.Get(); on {
		// Angle
		colour := termbox.ColorWhite
		if _, on := frame.Options.Option_Angle_Moving.Get(); on {
			if s.Accelerating {
				colour = termbox.ColorGreen
			} else {
				colour = termbox.ColorRed
			}
		}
		draw.PrintHorizontal(
			fill_width,         // width,
			fill_height,        // height,
			termbox.ColorBlack, // background,
			colour,             // foreground,
			x_from,             // x_from,
			y_from,             // y_from,
			fmt.Sprintf("%d°", int(movement.GetAngle())), // s,
		)
		y_from++
	}
	if _, on := frame.Options.Option_LatLong.Get(); on {
		// Lat/Long
		colour := termbox.ColorWhite
		if _, on := frame.Options.Option_LatLong_Moving.Get(); on {
			if s.Accelerating {
				colour = termbox.ColorGreen
			} else {
				colour = termbox.ColorRed
			}
		}
		lat, long := movement.GetLatLong()
		draw.PrintHorizontal(
			fill_width,         // width,
			fill_height,        // height,
			termbox.ColorBlack, // background,
			colour,             // foreground,
			x_from,             // x_from,
			y_from,             // y_from,
			fmt.Sprintf("%f, %f", lat, long), // s,
		)
		y_from++
	}
	if _, on := frame.Options.Option_Plugin_Offline.Get(); on && frame.Plugin != nil {
		// only display if plugin is enabled
		info := frame.Plugin.GetInfo()
		if info.LastSuccess.IsZero() {
			// we've never been online
			draw.PrintHorizontal(
				fill_width,         // width,
				fill_height,        // height,
				termbox.ColorBlack, // background,
				termbox.ColorRed,   // foreground,
				x_from,             // x_from,
				y_from,             // y_from,
				"Plugin Not Online Yet!",
			)
		} else {
			// we've been online
			if info.Offline {
				// we're offline
				draw.PrintHorizontal(
					fill_width,         // width,
					fill_height,        // height,
					termbox.ColorBlack, // background,
					termbox.ColorRed,   // foreground,
					x_from,             // x_from,
					y_from,             // y_from,
					fmt.Sprintf("Plugin Offline: %s", duration.Round(info.SinceLast, time.Second).String()), // s,
				)
			}
		}
	}
}
