package main

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sabey.co/spoofgo/cli"
	"sync"
)

var (
	location    int
	location_mu sync.RWMutex
)

func SetLocation(
	jump int,
) {
	location_mu.Lock()
	defer location_mu.Unlock()
	if jump != cli.PLAYER &&
		jump != cli.OPTIONS &&
		jump != cli.HELP &&
		jump != cli.NEWS {
		// invalid location
		return
	}
	// valid location
	location = jump
	if jump == cli.OPTIONS {
		// reset option position
		menu_option.Top()
	}
	if jump == cli.HELP {
		// reset help position
		textblock_help.Top()
	}
	if jump == cli.NEWS {
		// reset news textblock
		ResetNews()
	}
}
func GetLocation() int {
	location_mu.RLock()
	defer location_mu.RUnlock()
	return location
}
