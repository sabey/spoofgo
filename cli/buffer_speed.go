package cli

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"sabey.co/spoofgo/cli/draw-termbox"
	"sabey.co/spoofgo/movement"
)

func BufferSpeed(
	frame *Frame,
) []*draw.Buffer {
	speed := movement.GetSpeed()
	_, mph := frame.Options.Option_Speed_Unit.Get()
	if mph {
		// mp/h
		speed *= movement.SPEED_KMH_TO_MPH
	}
	s := fmt.Sprintf("%.2f", speed)           // 2 precision
	buff := make([]*draw.Buffer, 0, len(s)+5) // `.km/h`
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
	if mph {
		// mp/h
		buff = append(
			buff,
			&draw.Buffer{
				C:          'M',
				Background: termbox.ColorBlack,
				Foreground: termbox.ColorBlue,
			},
		)
	} else {
		// km/h
		buff = append(
			buff,
			&draw.Buffer{
				C:          'K',
				Background: termbox.ColorBlack,
				Foreground: termbox.ColorBlue,
			},
		)
	}
	return buff
}
