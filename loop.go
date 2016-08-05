package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"github.com/nsf/termbox-go"
	"log"
	"sabey.co/spoofgo/cli"
	"sabey.co/spoofgo/movement"
	"time"
)

func Loop() {
	// start
	if err := termbox.Init(); err != nil {
		// failed
		log.Fatalf("Failed to start CLI: \"%s\"\n", err)
		return
	}
	// done loading
	SetLocation(cli.PLAYER)
	// draw
	DrawCLI()
	// started, accept events
	event := make(chan termbox.Event)
	go func() {
		for {
			event <- termbox.PollEvent()
		}
	}()
	// loop and draw
loop:
	for {
		select {
		case e := <-event:
			// bug: if you're holding a key down and press another key, the held down key is ignored
			// so if we're holding forward and then turn, we stop registering forward
			// it would be nice if we could make things work so we could hold multiple keys at once
			// I don't know if it has to do with my terminal or not, but this should be looked at
			// todo: it would be nice to provide an option for strafing and backpedaling
			if e.Type == termbox.EventKey {
				loc := GetLocation()
				if loc == cli.PLAYER {
					if e.Key == termbox.KeyEsc {
						SetLocation(cli.OPTIONS)
					} else if e.Key == termbox.KeyF8 {
						SetLocation(cli.HELP)
					} else if e.Key == termbox.KeyF12 {
						SetLocation(cli.NEWS)
					} else if e.Ch == '`' ||
						e.Ch == '~' {
						// '-' and '+' were conflicting with `e.Key == termbox.KeyCtrlTilde` ???
						movement.ToggleMode()
					} else if e.Ch == '1' {
						movement.SetMode(movement.WALK)
					} else if e.Ch == '2' {
						movement.SetMode(movement.JOG)
					} else if e.Ch == '3' {
						movement.SetMode(movement.BICYCLE)
					} else if e.Ch == '4' {
						movement.SetMode(movement.CAR)
					} else if e.Ch == 'q' ||
						e.Ch == 'Q' {
						// top left
						movement.NorthWest()
					} else if e.Key == termbox.KeyArrowUp ||
						e.Ch == 'w' ||
						e.Ch == 'W' {
						// accelerate
						movement.Accelerate()
					} else if e.Ch == 'e' ||
						e.Ch == 'E' {
						// top right
						movement.NorthEast()
					} else if e.Key == termbox.KeyArrowLeft ||
						e.Ch == 'a' ||
						e.Ch == 'A' {
						// turn left
						movement.Left()
					} else if e.Key == termbox.KeyArrowDown ||
						e.Ch == 's' ||
						e.Ch == 'S' {
						// decelerate
						movement.Decelerate()
					} else if e.Key == termbox.KeyArrowRight ||
						e.Ch == 'd' ||
						e.Ch == 'D' {
						// turn right
						movement.Right()
					} else if e.Ch == 'z' ||
						e.Ch == 'Z' {
						// bottom left
						movement.SouthWest()
					} else if e.Ch == 'x' ||
						e.Ch == 'X' {
						// flip 180 degrees
						movement.Flip()
					} else if e.Ch == 'c' ||
						e.Ch == 'C' {
						// bottom right
						movement.SouthEast()
					} else if e.Ch == 'r' ||
						e.Ch == 'R' ||
						e.Ch == '=' ||
						e.Ch == '+' ||
						e.Key == termbox.KeyPgup { // Page Up
						movement.IncreaseModifier()
					} else if e.Ch == 'f' ||
						e.Ch == 'F' ||
						e.Ch == '-' ||
						e.Ch == '_' ||
						e.Key == termbox.KeyPgdn { // Page Down
						movement.DecreaseModifier()
					} else if e.Ch == 'v' ||
						e.Ch == 'V' ||
						e.Ch == '0' {
						movement.ResetModifier()
					}
				} else if loc == cli.OPTIONS {
					// OPTIONS
					if e.Key == termbox.KeyEsc {
						SetLocation(cli.PLAYER)
					} else if e.Key == termbox.KeyF8 {
						SetLocation(cli.HELP)
					} else if e.Key == termbox.KeyF12 {
						SetLocation(cli.NEWS)
					} else if e.Key == termbox.KeyEnter ||
						e.Key == termbox.KeySpace {
						OptionToggle()
					} else if e.Key == termbox.KeyArrowUp {
						menu_option.Up()
					} else if e.Key == termbox.KeyPgup {
						menu_option.JumpUp(5)
					} else if e.Key == termbox.KeyArrowDown {
						menu_option.Down()
					} else if e.Key == termbox.KeyPgdn {
						menu_option.JumpDown(5)
					} else if e.Key == termbox.KeyHome {
						menu_option.Top()
					} else if e.Key == termbox.KeyEnd {
						menu_option.Bottom()
					}
					// OPTIONS
				} else if loc == cli.HELP {
					// HELP
					if e.Key == termbox.KeyEsc {
						SetLocation(cli.OPTIONS)
					} else if e.Key == termbox.KeyF8 {
						SetLocation(cli.PLAYER)
					} else if e.Key == termbox.KeyF12 {
						SetLocation(cli.NEWS)
					} else if e.Key == termbox.KeyArrowUp {
						textblock_help.Up()
					} else if e.Key == termbox.KeyPgup {
						textblock_help.JumpUp(5)
					} else if e.Key == termbox.KeyArrowDown {
						textblock_help.Down()
					} else if e.Key == termbox.KeyPgdn {
						textblock_help.JumpDown(5)
					} else if e.Key == termbox.KeyHome {
						textblock_help.Top()
					} else if e.Key == termbox.KeyEnd {
						textblock_help.Bottom()
					}
					// HELP
				} else if loc == cli.NEWS {
					// NEWS
					if e.Key == termbox.KeyEsc {
						SetLocation(cli.OPTIONS)
					} else if e.Key == termbox.KeyF8 {
						SetLocation(cli.HELP)
					} else if e.Key == termbox.KeyF12 {
						SetLocation(cli.PLAYER)
					} else if e.Key == termbox.KeyArrowUp {
						news_mu.RLock()
						textblock_news.Up()
						news_mu.RUnlock()
					} else if e.Key == termbox.KeyPgup {
						news_mu.RLock()
						textblock_news.JumpUp(5)
						news_mu.RUnlock()
					} else if e.Key == termbox.KeyArrowDown {
						news_mu.RLock()
						textblock_news.Down()
						news_mu.RUnlock()
					} else if e.Key == termbox.KeyPgdn {
						news_mu.RLock()
						textblock_news.JumpDown(5)
						news_mu.RUnlock()
					} else if e.Key == termbox.KeyHome {
						news_mu.RLock()
						textblock_news.Top()
						news_mu.RUnlock()
					} else if e.Key == termbox.KeyEnd {
						news_mu.RLock()
						textblock_news.Bottom()
						news_mu.RUnlock()
					}
					// NEWS
				} else {
					// ??
					if e.Key == termbox.KeyEsc {
						// reset to player
						SetLocation(cli.PLAYER)
					}
				}
			}
			// currently we're drawing after keypress
			DrawCLI()
		case <-quit:
			break loop
		default:
			// we probably don't need to draw that often if nothing is changing
			DrawCLI()
			time.Sleep(50 * time.Millisecond)
		}
	}
	// close
	termbox.Close()
}
